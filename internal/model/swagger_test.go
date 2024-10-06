package model_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/difmaj/swaggo-gen/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMarshalJson(t *testing.T) {
	swaggerModel := new(model.Swagger)

	url := "https://developer.bling.com.br/build/assets/openapi-2aa7117f.json"

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to download JSON: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Received non-200 response status: %s", resp.Status)
	}

	fileContent, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	if err := json.Unmarshal(fileContent, &swaggerModel); err != nil {
		t.Fatalf("Error parsing ref: %v", err)
	}

	modelIndented, err := json.MarshalIndent(swaggerModel, "", "    ")
	if err != nil {
		t.Fatalf("Error indenting: %v", err)
	}

	var originalJson map[string]any
	var marshaledJson map[string]any

	if err := json.Unmarshal(fileContent, &originalJson); err != nil {
		t.Fatalf("Error unmarshaling original JSON: %v", err)
	}

	if err := json.Unmarshal(modelIndented, &marshaledJson); err != nil {
		t.Fatalf("Error unmarshaling marshaled JSON: %v", err)
	}

	assert.Equal(t, originalJson, marshaledJson, "The JSON structures are not equal.")
}
