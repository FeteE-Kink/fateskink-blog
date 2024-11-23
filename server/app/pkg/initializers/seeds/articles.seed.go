package seeds

import (
	"log"
	"math/rand"
	"server/app/database"
	"server/app/models"
	"server/app/pkg/helpers"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func ArticleSeed() {
	database := database.Db

	if !database.Migrator().HasTable(&models.Article{}) {
		return
	}

	var count int64
	database.Table("articles").Count(&count)

	if count > 0 {
		return
	}

	articles := []models.Article{}
	var user models.User
	result := database.Table("users").First(&user)

	if result.Error != nil {
		log.Fatalf("Error fetching user: %v", result.Error)
	}

	for i := 1; i <= 100; i++ {
		article := models.Article{
			Title:   gofakeit.Sentence(rand.Intn(10) + 6),
			Content: gofakeit.Paragraph(rand.Intn(10)+6, rand.Intn(10)+6, rand.Intn(100)+6, "\n\n"),
			Slug:    helpers.NewUUID(),
			UserId:  user.Id,
		}

		articles = append(articles, article)
	}

	database.Session(&gorm.Session{SkipHooks: true}).Create(&articles)
}
