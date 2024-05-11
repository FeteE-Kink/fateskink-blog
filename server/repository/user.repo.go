package repository

import (
	"context"
	"server/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(c *context.Context, db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// Find finds a user by their attribute.
func (r *UserRepository) Find(user *models.User) error {
	dbTables := r.db.Table("users")

	return dbTables.Where(&user).First(&user).Error
}
