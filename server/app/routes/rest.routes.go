package routes

import (
	"net/http"
	"server/app/controllers"
	rest "server/app/routes/rest"

	"github.com/gin-gonic/gin"
)

func RegisterRestRoutes(r *gin.Engine) {
	r.GET("healthz", healthCheck)
	r.GET("/", controllers.HomeController)
	r.GET("/signin", controllers.SignInController)

	rest.UploadRoutes(r)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
