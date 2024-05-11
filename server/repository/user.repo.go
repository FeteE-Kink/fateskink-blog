package repository

import (
	"context"
	"errors"
	"server/models"

	"golang.org/x/crypto/bcrypt"
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

func (r *UserRepository) Auth(email string, password string) (user *models.User, err error) {
	u := models.User{Email: email}

	userFindErr := r.Find(&u)
	if userFindErr != nil {
		return nil, errors.New("can not find user")
	}

	comparePwError := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	if comparePwError != nil {
		return nil, errors.New("email or password is incorrect")
	}

	return &u, nil
}
