package globalTypes

import (
	"server/app/models"
	"server/app/pkg/helpers"
)

type MetadataPayloadType struct {
	Metadata *models.Metadata
}

func (mt *MetadataPayloadType) Total() *int32 {
	return helpers.Int32Pointer(mt.Metadata.Total)
}

func (mt *MetadataPayloadType) Page() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Page))
}

func (mt *MetadataPayloadType) Pages() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Pages))
}

func (mt *MetadataPayloadType) PerPage() *int32 {
	return helpers.Int32Poiter(int32(mt.Metadata.PerPage))
}

func (mt *MetadataPayloadType) Count() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Count))
}

func (mt *MetadataPayloadType) Next() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Next))
}

func (mt *MetadataPayloadType) Prev() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Prev))
}

func (mt *MetadataPayloadType) From() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.From))
}

func (mt *MetadataPayloadType) To() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.To))
}
