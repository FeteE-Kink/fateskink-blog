package gqlInput

import "server/app/models"

type PagyInput struct {
	PerPage *int32
	Page    *int32
}

func (input *PagyInput) ToPaginatioInput() models.PaginationData {
	paginationInput := models.PaginationData{
		Metadata: models.Metadata{
			Page:    1,
			PerPage: 10,
		},
	}

	if input != nil {
		if input.Page != nil && *input.Page >= 1 {
			paginationInput.Metadata.Page = int(*input.Page)
		}

		if input.PerPage != nil && *input.PerPage >= 1 {
			paginationInput.Metadata.PerPage = int(*input.PerPage)
		}
	}

	return paginationInput
}
