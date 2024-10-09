package parsing_test

import (
	"os"
	"testing"

	"github.com/difmaj/swaggo-gen/internal/parsing"
	"github.com/difmaj/swaggo-gen/internal/test"
	"github.com/stretchr/testify/suite"
)

type ParsingSuite struct {
	suite.Suite
	dir string
}

func (rs *ParsingSuite) SetupSuite() {
	rs.dir = "tmp/schemas"
}

func (rs *ParsingSuite) TearDownSuite() {
	os.RemoveAll(rs.dir)
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(ParsingSuite))
}

func (rs *ParsingSuite) TestGenerateSchemas() {
	outModel, err := test.GetModel()
	if err != nil {
		rs.T().Fatalf("Error on getting model: %v", err)
	}

	parsing.GenerateSchemas(outModel.SwaggerModel, "tmp")

	rs.DirExists(rs.dir)
}
