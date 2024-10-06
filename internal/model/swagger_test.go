package model_test

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/difmaj/swaggo-gen/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMarshalJson(t *testing.T) {
	swaggerModel := new(model.Swagger)

	file, err := os.Open("../../examples/swagger-bling.json")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if err := json.Unmarshal(fileContent, &swaggerModel); err != nil {
		t.Fatalf("Error parsing ref: %v", err)
	}

	modelIndented, err := json.MarshalIndent(swaggerModel, "", "    ")
	if err != nil {
		t.Fatalf("Error indenting: %v", err)
	}

	var originalJson map[string]interface{}
	var marshaledJson map[string]interface{}

	if err := json.Unmarshal(fileContent, &originalJson); err != nil {
		t.Fatalf("Error unmarshaling original JSON: %v", err)
	}

	if err := json.Unmarshal(modelIndented, &marshaledJson); err != nil {
		t.Fatalf("Error unmarshaling marshaled JSON: %v", err)
	}

	assert.Equal(t, originalJson, marshaledJson, "The JSON structures are not equal.")
}
