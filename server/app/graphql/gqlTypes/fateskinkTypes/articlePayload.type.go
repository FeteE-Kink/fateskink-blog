package fateskinktypes

import (
	"fmt"
	"server/app/models"
	"strings"

	"github.com/graph-gophers/graphql-go"
)

type ArticlePayloadType struct {
	Article *models.Article
}

func (a *ArticlePayloadType) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(a.Article.Id))
}

func (a *ArticlePayloadType) Title() string {
	return a.Article.Title
}

func (a *ArticlePayloadType) Content() string {
	return a.Article.Content
}

func (a *ArticlePayloadType) PreviewContent() string {
	if a.Article.Content == "" {
		return ""
	}

	words := strings.Split(a.Article.Content, " ")

	if len(words) > 100 {
		words = words[:100]
	}

	return strings.Join(words, " ")
}
