package seeds

import (
	"log"
	"server/app/database"
	"server/app/models"

	"gorm.io/gorm"
)

func TagSeed() {
	database := database.Db

	if !database.Migrator().HasTable(&models.Tag{}) {
		return
	}

	var count int64
	database.Table("tags").Count(&count)

	if count > 0 {
		return
	}

	tags := []models.Tag{}

	listTags := []string{
		"All",
		"DevOps",
		"Amazon Web Services",
		"Distributed Systems",
		"Linux",
		"Web Development",
		"System Design",
		"AWS Web Services",
		"Terraform",
		"CSS",
		"Psychology",
		"Kubernetes",
		"Cheatsheet",
		"Data Structures",
		"Scalability",
		"Ansible",
		"Docker",
		"Git",
		"Accessibility",
		"Android",
		"Career",
	}

	var user models.User
	result := database.Table("users").First(&user)

	if result.Error != nil {
		log.Fatalf("Error fetching user: %v", result.Error)
	}

	for _, name := range listTags {
		tags = append(tags, models.Tag{
			Name:   name,
			UserId: user.Id,
		})
	}

	database.Session(&gorm.Session{SkipHooks: true}).Create(&tags)
}
