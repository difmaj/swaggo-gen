package parsing

import (
	"fmt"

	"github.com/difmaj/swaggo-gen/internal/model"
	"github.com/difmaj/swaggo-gen/internal/utils"
)

func GenerateSchemas(swaggerModel *model.Swagger, outputPath string) {
	err := utils.CreatePath(fmt.Sprintf("%s/schemas", outputPath))
	if err != nil {
		return
	}

	fmt.Println("Directory created successfully:", outputPath)
}
