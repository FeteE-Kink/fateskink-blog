package services

import (
	"context"
	"server/app/exceptions"
	"server/app/models"
	"server/app/pkg/helpers"
	"server/app/repository"

	"gorm.io/gorm"
)

type AuthService struct {
	Email    string
	Password string
	Ctx      *context.Context
	Db       *gorm.DB

	Token *string
}

func (a *AuthService) Execute() error {
	if err := a.validate(); err != nil {
		return err
	}
	user, err := a.findUser()
	if err != nil {
		return err
	}

	token, _ := helpers.GenerateJwtToken(user.GenerateJwtClaims())
	a.Token = &token
	return nil
}

func (a *AuthService) validate() error {
	exception := exceptions.NewUnprocessableContentError("", nil)
	if a.Email == "" {
		exception.AddError("email", []interface{}{"can not be empty"})
	}

	if a.Password == "" {
		exception.AddError("password", []interface{}{"can not be empty"})
	}
	if len(exception.Errors) > 0 {
		return exception
	}

	return nil
}

func (a *AuthService) findUser() (*models.User, error) {
	repo := repository.NewUserRepository(a.Ctx, a.Db)
	user, err := repo.Auth(a.Email, a.Password)

	if err != nil {
		return nil, exceptions.NewUnprocessableContentError("base", exceptions.ResourceModificationError{
			"password": {"User Or Password is incorrect"},
		})
	}

	return user, nil
}
