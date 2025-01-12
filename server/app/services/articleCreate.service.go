package services

import (
	"server/app/exceptions"
	fateskinkInputs "server/app/graphql/inputs/feateskinkInputs"
	"server/app/models"

	"gorm.io/gorm"
)

type ArticleCreateService struct {
	Input *fateskinkInputs.ArticleCreateInput
	User  *models.User
	Db    *gorm.DB
}

func (service ArticleCreateService) Execute() error {
	if service.Input == nil {
		return exceptions.NewUnprocessableContentError("Input Incorrect", exceptions.ResourceModificationError{
			"base": []interface{}{"Input Incorrect"},
		})
	}

	// form := forms.New
	return nil
}
