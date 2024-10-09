package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/difmaj/swaggo-gen/internal/model"
)

type GetModelOutput struct {
	SwaggerModel *model.Swagger
	FileContent  []byte
}

func GetModel() (*GetModelOutput, error) {
	swaggerModel := new(model.Swagger)

	url := "https://developer.bling.com.br/build/assets/openapi-2aa7117f.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %s", resp.Status)
	}

	fileContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(fileContent, &swaggerModel); err != nil {
		return nil, err
	}

	return &GetModelOutput{
		SwaggerModel: swaggerModel,
		FileContent:  fileContent,
	}, nil
}
