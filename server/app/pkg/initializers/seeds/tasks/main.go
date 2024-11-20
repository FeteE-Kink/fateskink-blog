package main

import (
	"os"
	"server/app/pkg/initializers"
	"server/app/pkg/initializers/seeds"
)

func main() {
	initializers.LoadEnv()
	initializers.LoadDb()

	SeedCore()

	if os.Getenv("APP_ENV") == "development" {
		SeedDev()
	}
}

func SeedCore() {
	seeds.UserSeed()
}

func SeedDev() {
	seeds.TagSeed()
	seeds.ArticleSeed()
}
