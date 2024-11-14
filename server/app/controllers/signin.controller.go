package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInController(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}
