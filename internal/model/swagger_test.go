package model_test

import (
	"encoding/json"
	"testing"

	"github.com/difmaj/swaggo-gen/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestMarshalJson(t *testing.T) {
	outModel, err := test.GetModel()
	if err != nil {
		t.Fatalf("Error on getting model: %v", err)
	}

	modelIndented, err := json.MarshalIndent(outModel.SwaggerModel, "", "    ")
	if err != nil {
		t.Fatalf("Error indenting: %v", err)
	}

	var originalJson map[string]any
	var marshaledJson map[string]any

	if err := json.Unmarshal(outModel.FileContent, &originalJson); err != nil {
		t.Fatalf("Error unmarshaling original JSON: %v", err)
	}

	if err := json.Unmarshal(modelIndented, &marshaledJson); err != nil {
		t.Fatalf("Error unmarshaling marshaled JSON: %v", err)
	}

	assert.Equal(t, originalJson, marshaledJson, "The JSON structures are not equal.")
}
