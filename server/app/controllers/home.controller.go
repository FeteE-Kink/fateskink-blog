package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Title    string
	Metadata []Metadata
}

type Metadata struct {
	Name    string
	Content string
}

func HomeController(c *gin.Context) {
	data := PageData{
		Title: "Fateskink",
		Metadata: []Metadata{
			{"description", "This is home page"},
			{"keywords", "home, page, article"},
		},
	}

	c.HTML(http.StatusOK, "index.html", data)
}
