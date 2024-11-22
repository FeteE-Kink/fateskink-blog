package controllers

import "server/app/pkg/helpers"

type PageData struct {
	BaseUrl  string
	Title    string
	Metadata []Metadata
}

type Metadata struct {
	Name    string
	Content string
}

func NewMetadata(title *string, metadata *[]Metadata) PageData {
	baseUrl := helpers.GetEnv("BASE_URL", "http://localhost:3007/")
	pageTitle := "Fateskink"
	if title != nil {
		pageTitle = *title
	}

	pageMetadata := []Metadata{}
	if metadata == nil {
		pageMetadata = []Metadata{
			{
				Name:    "",
				Content: "",
			},
		}
	}
	return PageData{
		BaseUrl:  baseUrl,
		Title:    pageTitle,
		Metadata: pageMetadata,
	}
}
