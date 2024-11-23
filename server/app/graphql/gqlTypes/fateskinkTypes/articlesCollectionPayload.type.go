package fateskinktypes

import globaltypes "server/app/graphql/gqlTypes/globalTypes"

type ArticlesCollectionPayloadType struct {
	Collection *[]*ArticlePayloadType
	Metadata   *globaltypes.MetadataPayloadType
}
