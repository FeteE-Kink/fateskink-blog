package initializers

import (
	"log"
	"server/app/database"
	"server/app/pkg/helpers"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() {
	appEnv := helpers.GetEnv("APP_ENV", "development")
	if appEnv == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Can not load .env file: ", err)
		}
	}
}

func LoadDb() {
	database.InitDb()
}
