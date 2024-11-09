package setup

import "server/app/pkg/initializers"

func InitEnvironment() {
	initializers.LoadEnv()
	initializers.InitLogger()
	initializers.LoadDb()
}
