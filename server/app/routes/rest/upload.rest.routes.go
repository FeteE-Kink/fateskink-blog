package routes

import (
	"server/app/controllers"

	"github.com/gin-gonic/gin"
)

func UploadRoutes(r *gin.Engine) {
	r.POST("fateskinkUploads", controllers.UploadHandler)
}
