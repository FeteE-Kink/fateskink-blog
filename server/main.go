package main

import (
	"os"
	"server/setup"
)

func main() {
	os.Setenv("TZ", "UTC")

	setup.InitEnvironment()

	server := setup.InitServer()

	go setup.StartServer(server)
	setup.WaitForShutDown(server)
}
