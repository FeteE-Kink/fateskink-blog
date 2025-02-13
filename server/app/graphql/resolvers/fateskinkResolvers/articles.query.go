package fateskinkResolvers

import (
	"context"
	"server/app/graphql/gqlTypes/fateskinkTypes"
	fateskinktypes "server/app/graphql/gqlTypes/fateskinkTypes"
	"server/app/graphql/gqlTypes/globalTypes"
	"server/app/graphql/inputs/globalInputs"
	"server/app/models"
	"server/app/repository"
)

func (r *Resolver) Articles(
	ctx context.Context, args globalInputs.ArtilcesInput,
) (
	*fateskinkTypes.ArticlesCollectionPayloadType, error,
) {
	var articles []*models.Article
	paginationData, articlesQuery := args.ToPaginationDatAndQuery()

	repo := repository.NewArticleRepository(r.Db)
	if err := repo.List(&articles, &paginationData, &articlesQuery); err != nil {
		return nil, err
	}

	return &fateskinktypes.ArticlesCollectionPayloadType{
		Collection: sliceToTypes(articles),
		Metadata: &globalTypes.MetadataPayloadType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}

func sliceToTypes(articles []*models.Article) *[]*fateskinktypes.ArticlePayloadType {
	resolvers := make([]*fateskinktypes.ArticlePayloadType, len(articles))
	for i, e := range articles {
		resolvers[i] = &fateskinktypes.ArticlePayloadType{Article: e}
	}

	return &resolvers
}
