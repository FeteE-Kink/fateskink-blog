package initializers

import (
	"log"
	"os"
	"server/app/graphql/resolvers/fateskinkResolvers"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"gorm.io/gorm"
)

func FateskinkGqlHandler(db *gorm.DB) gin.HandlerFunc {
	schema, err := fetchSchema("./app/graphql/schemas/fateskink/")

	if err != nil {
		log.Fatalf("failed to get schema: %v", err)
	}

	opts := []graphql.SchemaOpt{graphql.UseStringDescriptions(), graphql.UseFieldResolvers()}
	gqlSchema := graphql.MustParseSchema(schema, &fateskinkResolvers.Resolver{Db: db}, opts...)

	return ginSchemaHandler(schema, db, gqlSchema)
}

func fetchSchema(schemaPath string) (string, error) {
	entries, err := os.ReadDir(schemaPath)
	if err != nil {
		return "", err
	}

	var schemaContent []byte
	for _, entry := range entries {
		filePath := schemaPath + entry.Name()
		content, err := os.ReadFile(filePath)

		if err != nil {
			log.Printf("failed to read file %s: %v", filePath, err)
			continue
		}

		schemaContent = append(schemaContent, content...)
	}
	return string(schemaContent), nil
}

func ginSchemaHandler(schema string, db *gorm.DB, gqlSchema *graphql.Schema) gin.HandlerFunc {
	handler := &relay.Handler{Schema: gqlSchema}

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
