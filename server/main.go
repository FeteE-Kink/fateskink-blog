package main

import (
	"os"
	"server/pkg/initializers"
	"server/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	initializers.LoadEnv()

	// db := database.InitDb()

	r := gin.Default()
	r.Use(initializers.CorsConfig())
	r.Use(logger.Logger(logrus.New()), gin.Recovery())
	// r.POST("/fateskinkGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.InsightGqlHandler(db))
}
