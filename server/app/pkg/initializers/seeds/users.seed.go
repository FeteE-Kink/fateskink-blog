package seeds

import (
	"server/app/database"
	"server/app/models"

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
		Email:             "admin@example.com",
		EncryptedPassword: "$2a$10$8080LRcosHOEqh6kpHbQie/vgubdGQkXIvPQUkAaCzGZCs6FMsmc6",
		Name:              "Admin",
		FullName:          "Fateskink",
	}

	database.Session(&gorm.Session{SkipHooks: true}).Create(&user)
}
