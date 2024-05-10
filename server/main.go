package main

import (
	"fmt"
	"os"
	"server/database"
	"server/pkg/initializers"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	initializers.LoadEnv()

	db := database.InitDb()
	fmt.Print(db)
}
