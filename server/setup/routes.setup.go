package setup

import (
	"html/template"
	"os"
	"path/filepath"
	"server/app/pkg/initializers"
	"server/app/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(initializers.CorsConfig())
	r.Use(initializers.LoggingMiddleware())

	setupHTMLRouter(r)
	setupAPIRouter(r)

	return r
}

func setupAPIRouter(r *gin.Engine) {
	routes.RegisterGraphQLRouter(r)
	routes.RegisterRestRoutes(r)
}

func setupHTMLRouter(r *gin.Engine) {
	templates, err := collectTemplates("app/templates")
	if err != nil {
		log.Errorf("Failed to collect templates: %+v", err)
	}

	tmpl := template.Must(template.ParseFiles(templates...))
	r.SetHTMLTemplate(tmpl)

	r.Static("/static", "./app/static")

	if os.Getenv("APP_ENV") == "development" {
		r.Static("/uploads", "./app/tmp/uploads")
	}
}

func collectTemplates(root string) ([]string, error) {
	var templates []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			templates = append(templates, path)
		}

		return nil
	})

	return templates, err
}
