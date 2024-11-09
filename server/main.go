package main

import (
	"os"
	"server/setup"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	setup.InitEnvironment()
	server := setup.InitServer()

	go setup.StartServer(server)
	setup.WaitForShutDown(server)
}
