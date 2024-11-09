package controllers

import "github.com/gin-gonic/gin"

func UploadHandler(c *gin.Context) {
	auth := AuthUser(c)
	if !auth {
		return
	}

	handleUpload(c)
}

func handleUpload(c *gin.Context) {}
