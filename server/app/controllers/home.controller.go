package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	title := "Fateskink"
	metadata := []Metadata{
		{"description", "This is home page"},
		{"keywords", "home, page, article"},
	}
	data := NewMetadata(&title, &metadata)

	c.HTML(http.StatusOK, "index.html", data)
}
