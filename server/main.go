package main

import (
	"os"
	"server/app/database"
	"server/app/pkg/auths"
	"server/app/pkg/initializers"
	"server/app/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	initializers.LoadEnv()

	db := database.InitDb()
	r := gin.Default()

	r.Use(initializers.CorsConfig())
	r.Use(logger.Logger(logrus.New()), gin.Recovery())

	setupHTMLRouter(r)
	r.POST("/fateskinkGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.FateskinkGqlHandler(db))

	r.Run()
}

func setupHTMLRouter(r *gin.Engine) {

}
