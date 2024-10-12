package controller

import (
	"net/http"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

// OptionFunc can be used to customize a new API client.
type OptionFunc func(*Controller) error

// WithCustomBackoff can be used to configure a custom backoff policy.
func WithCustomBackoff(backoff retryablehttp.Backoff) OptionFunc {
	return func(c *Controller) error {
		c.client.Backoff = backoff
		return nil
	}
}

// WithCustomLeveledLogger can be used to configure a custom retryablehttp
// leveled logger.
func WithCustomLeveledLogger(leveledLogger retryablehttp.LeveledLogger) OptionFunc {
	return func(c *Controller) error {
		c.client.Logger = leveledLogger
		return nil
	}
}

// WithCustomLimiter injects a custom rate limiter to the client.
func WithCustomLimiter(limiter RateLimiter) OptionFunc {
	return func(c *Controller) error {
		c.configureLimiterOnce.Do(func() {})
		c.limiter = limiter
		return nil
	}
}

// WithCustomLogger can be used to configure a custom retryablehttp logger.
func WithCustomLogger(logger retryablehttp.Logger) OptionFunc {
	return func(c *Controller) error {
		c.client.Logger = logger
		return nil
	}
}

// WithCustomRetry can be used to configure a custom retry policy.
func WithCustomRetry(checkRetry retryablehttp.CheckRetry) OptionFunc {
	return func(c *Controller) error {
		c.client.CheckRetry = checkRetry
		return nil
	}
}

// WithCustomRetryMax can be used to configure a custom maximum number of retries.
func WithCustomRetryMax(retryMax int) OptionFunc {
	return func(c *Controller) error {
		c.client.RetryMax = retryMax
		return nil
	}
}

// WithCustomRetryWaitMinMax can be used to configure a custom minimum and
// maximum time to wait between retries.
func WithCustomRetryWaitMinMax(waitMin, waitMax time.Duration) OptionFunc {
	return func(c *Controller) error {
		c.client.RetryWaitMin = waitMin
		c.client.RetryWaitMax = waitMax
		return nil
	}
}

// WithErrorHandler can be used to configure a custom error handler.
func WithErrorHandler(handler func(*http.Response) error) OptionFunc {
	return func(c *Controller) error {
		c.errorHandler = handler
		return nil
	}
}

// WithHTTPClient can be used to configure a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) OptionFunc {
	return func(c *Controller) error {
		c.client.HTTPClient = httpClient
		return nil
	}
}

// WithoutRetries disables the default retry logic.
func WithoutRetries() OptionFunc {
	return func(c *Controller) error {
		c.disableRetries = true
		return nil
	}
}
