package main

import (
	"os"
	"server/setup"
)

func main() {
	os.Setenv("TZ", "UTC")

	setup.InitEnvironment()

	// r := gin.Default()

	// r.POST("/fateskinkGql",
	// 	auths.GinContextToContextMiddleware(),
	// 	auths.JwtTokenCheck,
	// 	initializers.FateskinkGqlHandler(database.Db),
	// )

	// r.Run()
	server := setup.InitServer()

	go setup.StartServer(server)
	setup.WaitForShutDown(server)
}
