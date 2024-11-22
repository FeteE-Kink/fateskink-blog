package seeds

import (
	"server/app/database"
	"server/app/models"
	"server/app/pkg/helpers"

	"gorm.io/gorm"
)

func UserSeed() {
	database := database.Db

	if !database.Migrator().HasTable(&models.User{}) {
		return
	}

	var count int64
	database.Table("users").Count(&count)

	if count > 0 {
		return
	}

	// password: 12341234
	user := models.User{
		Email:             helpers.GetEnv("ADMIN_EMAIL", "admin@example.com"),
		EncryptedPassword: "$2a$10$8080LRcosHOEqh6kpHbQie/vgubdGQkXIvPQUkAaCzGZCs6FMsmc6",
		Name:              "Admin",
		FullName:          "Admin",
	}

	database.Session(&gorm.Session{SkipHooks: true}).Create(&user)
}
