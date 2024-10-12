package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-cleanhttp"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"golang.org/x/time/rate"
)

const (
	baseURLFormat   = "https://%s"
	headerRateLimit = "X-RateLimit-Limit"
	headerRateReset = "X-RateLimit-Reset"
)

// A Controller manages communication with the API.
type Controller struct {
	// HTTP client used to communicate with the API.
	client *retryablehttp.Client
	// Base URL for API requests.
	baseURL *url.URL
	// disableRetries is used to disable the default retry logic.
	disableRetries bool
	// configureLimiterOnce is used to configure the rate limiter only once.
	configureLimiterOnce sync.Once
	// Limiter is used to limit API calls and prevent 429 responses.
	limiter RateLimiter
	// errorHandler is a function that checks the response for errors.
	errorHandler func(*http.Response) error
}

// RateLimiter describes the interface that all (custom) rate limiters must implement.
type RateLimiter interface {
	Wait(context.Context) error
}

// NewController returns a new API controller.
func NewController(baseURL string, options ...OptionFunc) (*Controller, error) {
	controller, err := newController(options...)
	if err != nil {
		return nil, err
	}

	// Set the default base URL.
	controller.setBaseURL(baseURL)
	return controller, nil
}

func newController(options ...OptionFunc) (*Controller, error) {
	c := &Controller{}

	// Configure the HTTP client.
	c.client = &retryablehttp.Client{
		Backoff:      c.retryHTTPBackoff,
		CheckRetry:   c.retryHTTPCheck,
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		RetryMax:     5,
	}

	// Apply any given client options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	if c.limiter == nil {
		c.limiter = rate.NewLimiter(rate.Inf, 0)
	}
	return c, nil
}

// retryHTTPCheck provides a callback for Client.CheckRetry which
// will retry both rate limit (429) and server (>= 500) errors.
func (c *Controller) retryHTTPCheck(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	if err != nil {
		return false, err
	}
	if !c.disableRetries && (resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= http.StatusInternalServerError) {
		return true, nil
	}
	return false, nil
}

// retryHTTPBackoff provides a generic callback for Client.Backoff which
// will pass through all calls based on the status code of the response.
func (c *Controller) retryHTTPBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// Use the rate limit backoff function when we are rate limited.
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		return rateLimitBackoff(min, max, attemptNum, resp)
	}

	// Set custom duration's when we experience a service interruption.
	min = 700 * time.Millisecond
	max = 900 * time.Millisecond

	return retryablehttp.LinearJitterBackoff(min, max, attemptNum, resp)
}

// rateLimitBackoff provides a callback for Client.Backoff
func rateLimitBackoff(min, max time.Duration, _ int, resp *http.Response) time.Duration {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// First create some jitter bounded by the min and max durations.
	jitter := time.Duration(rnd.Float64() * float64(max-min))

	if resp != nil {
		if v := resp.Header.Get(headerRateReset); v != "" {
			if reset, _ := strconv.ParseInt(v, 10, 64); reset > 0 {
				// Only update min if the given time to wait is longer.
				if wait := time.Until(time.Unix(reset, 0)); wait > min {
					min = wait
				}
			}
		}
	}

	return min + jitter
}

// configureLimiter configures the rate limiter.
func (c *Controller) configureLimiter(ctx context.Context, headers http.Header) {
	if v := headers.Get(headerRateLimit); v != "" {
		if rateLimit, _ := strconv.ParseFloat(v, 64); rateLimit > 0 {

			rateLimit /= 60

			limit := rate.Limit(rateLimit * 0.66)
			burst := int(rateLimit * 0.33)

			c.limiter = rate.NewLimiter(limit, burst)
			c.limiter.Wait(ctx)
		}
	}
}

// BaseURL return a copy of the baseURL.
func (c *Controller) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Controller) setBaseURL(urlStr string) error {
	if !strings.Contains(urlStr, "://") {
		urlStr = fmt.Sprintf(baseURLFormat, urlStr)
	}

	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Update the base URL of the client.
	c.baseURL = baseURL
	return nil
}

// NewRequest creates a new API request.
func (c *Controller) NewRequest(ctx context.Context, method, path string, opt any) (*retryablehttp.Request, error) {

	u := c.BaseURL()
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	headers := make(http.Header)

	// Add the query string if any options were given.
	var body interface{}
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		headers.Set("Accept", "application/json")
		headers.Set("Content-Type", "application/json")

		if opt != nil {
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			}
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()

		// Create a request specific headers map.
		headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}

	// Create a new request using the given method and URL.
	req, err := retryablehttp.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range headers {
		req.Header[k] = v
	}
	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Controller) Do(req *retryablehttp.Request, v interface{}) (*http.Response, error) {

	// Wait will block until the limiter can obtain a new token.
	err := c.limiter.Wait(req.Context())
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)

	// If not yet configured, try to configure the rate limiter
	// using the response headers we just received. Fail silently
	// so the limiter will remain disabled in case of an error.
	c.configureLimiterOnce.Do(func() { c.configureLimiter(req.Context(), resp.Header) })

	if err := c.checkResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}

// checkResponse checks the API response for errors
func (c *Controller) checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent, http.StatusNotModified:
		return nil
	}
	return c.errorHandler(r)
}

// CloseIdleConnections closes the idle connections.
func (c *Controller) Shutdown() error {
	c.client.HTTPClient.CloseIdleConnections()
	return nil
}
