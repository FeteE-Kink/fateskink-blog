package repository

import (
	"fmt"
	"server/app/graphql/inputs/globalInputs"
	"server/app/models"
	"server/app/pkg/pagination"
	"strings"

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
	query *globalInputs.ArticlesQueryInput,
) error {
	dbTables := r.db.Model(&models.Article{})

	return dbTables.Scopes(
		pagination.Paginate(
			dbTables.Scopes(
				r.titleCount(query.TitleCont),
				r.slugCount(query.SlugCont),
				r.createdAtGteq(query.CreatedAtGteq),
				r.createdAtLteq(query.CreatedAtLteq),
			),
			paginationData, r.TableName),
	).Order("id desc").Find(&articles).Error
}

func (r *ArticleRepository) titleCount(titleCont *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleCont == nil {
			return db
		}

		return db.Where(gorm.Expr(
			`lower(title) like ?`,
			fmt.Sprintf("%%%s%%", strings.ToLower(*titleCont))),
		)
	}
}

func (r *ArticleRepository) slugCount(slugCont *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if slugCont == nil {
			return db
		}

		return db.Where(gorm.Expr(
			`lower(slug) like ?`,
			fmt.Sprintf("%%%s%%", strings.ToLower(*slugCont))),
		)
	}
}

func (r *ArticleRepository) createdAtGteq(createdAtGteq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if createdAtGteq == nil {
			return db
		}

		return db.Where("created_at >= ?", *createdAtGteq)
	}
}

func (r *ArticleRepository) createdAtLteq(createdAtLteq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if createdAtLteq == nil {
			return db
		}

		return db.Where("created_at >= ?", *createdAtLteq)
	}
}

func (r *ArticleRepository) Find(article *models.Article) error {
	dbTables := r.db.Model(&models.Article{})

	return dbTables.Where(&article).First(&article).Error
}

// func (r *ArticleRepository) tagIn(tagIn *[]int32) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if tagIn == nil {
// 			return db
// 		}

// 		return db.Where(gorm.Expr(
// 			`lower(slug) like ?`,
// 			fmt.Sprintf("%%%s%%", strings.ToLower(*slugCont))),
// 		)
// 	}
// }
