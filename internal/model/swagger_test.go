package model_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/difmaj/swaggo-gen/internal/model"
)

func TestMarshalJson(t *testing.T) {
	response := new(model.Swagger)

	file, err := os.Open("../../examples/swagger-bling.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	jsonRef, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(jsonRef, &response); err != nil {
		fmt.Println("Error parsing ref:", err)
	}

	fmt.Println(response)

	jsonStr, err := json.MarshalIndent(response, " ", "\t")

	if err != nil {
		fmt.Println("Error indenting:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonStr))
}
