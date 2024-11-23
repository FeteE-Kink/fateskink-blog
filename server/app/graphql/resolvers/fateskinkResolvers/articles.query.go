package fateskinkResolvers

import (
	"context"
	fateskinktypes "server/app/graphql/gqlTypes/fateskinkTypes"
	"server/app/graphql/gqlTypes/globalTypes"
	gqlInput "server/app/graphql/inputs"
	"server/app/models"
	"server/app/repository"
)

func (r *Resolver) Articles(
	ctx context.Context, args gqlInput.ArtilcesInput,
) (
	*fateskinktypes.ArticlesCollectionPayloadType, error,
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
