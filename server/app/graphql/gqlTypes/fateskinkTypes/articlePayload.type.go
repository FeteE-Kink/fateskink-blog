package fateskinktypes

import (
	"fmt"
	"server/app/models"

	"github.com/graph-gophers/graphql-go"
)

type ArticlePayloadType struct {
	Article *models.Article
}

func (a *ArticlePayloadType) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(a.Article.Id))
}
