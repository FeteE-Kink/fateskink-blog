package globalInputs

import (
	"server/app/models"
	"server/app/pkg/helpers"
	"strings"

	"github.com/graph-gophers/graphql-go"
)

type ArticleInput struct {
	Id graphql.ID
}

type ArtilcesInput struct {
	Input *PagyInput
	Query *ArticlesQueryInput
}

type ArticlesQueryInput struct {
	TitleCont     *string
	SlugCont      *string
	TagIn         *[]*int32
	CreatedAtGteq *string
	CreatedAtLteq *string
}

func (input *ArtilcesInput) ToPaginationDatAndQuery() (models.PaginationData, ArticlesQueryInput) {
	paginatioData := input.Input.ToPaginatioInput()
	query := ArticlesQueryInput{}

	if input.Query != nil {
		if input.Query.TitleCont != nil && strings.TrimSpace(*input.Query.TitleCont) != "" {
			query.TitleCont = input.Query.TitleCont
		}

		if input.Query.SlugCont != nil && strings.TrimSpace(*input.Query.SlugCont) != "" {
			query.SlugCont = input.Query.SlugCont
		}

		if input.Query.TagIn != nil && len(*input.Query.TagIn) > 0 {
			query.TagIn = input.Query.TagIn
		}

		if input.Query.CreatedAtGteq != nil && strings.TrimSpace(*input.Query.CreatedAtGteq) != "" {
			if parseDate, err := helpers.ParseStringToStartsOfDay(*input.Query.CreatedAtGteq); err != nil {
				query.CreatedAtGteq = parseDate
			}
		}

		if input.Query.CreatedAtLteq != nil && strings.TrimSpace(*input.Query.CreatedAtLteq) != "" {
			if parseDate, err := helpers.ParseStringToStartsOfDay(*input.Query.CreatedAtLteq); err != nil {
				query.CreatedAtLteq = parseDate
			}
		}
	}

	return paginatioData, query
}
