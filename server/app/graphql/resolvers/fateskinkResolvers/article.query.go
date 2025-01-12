package fateskinkResolvers

import (
	"context"
	"server/app/exceptions"
	fateskinkTypes "server/app/graphql/gqlTypes/fateskinkTypes"
	fateskinktypes "server/app/graphql/gqlTypes/fateskinkTypes"
	"server/app/models"
	"server/app/pkg/helpers"
	"server/app/repository"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

func (r *Resolver) Article(
	ctx context.Context, args struct{ ID graphql.ID },
) (*fateskinkTypes.ArticlePayloadType, error) {
	if args.ID == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	articleId, err := helpers.GqlIdtoInt32(args.ID)
	if err != nil {
		return nil, err
	}
	article := models.Article{Id: articleId}
	repo := repository.NewArticleRepository(r.Db)
	err = repo.Find(&article)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
	}

	return &fateskinktypes.ArticlePayloadType{Article: &article}, nil
}
