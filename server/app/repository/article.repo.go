package repository

import (
	gqlInput "server/app/graphql/inputs"
	"server/app/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	Repository
	TableName string
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		Repository: Repository{db: db},
		TableName:  "articles",
	}
}

func (r *ArticleRepository) List(
	articles *[]*models.Article,
	paginationData *models.PaginationData,
	query *gqlInput.ArticlesQueryInput,
) error {
	return nil
}
