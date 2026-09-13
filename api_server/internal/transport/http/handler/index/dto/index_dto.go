package dto

import "api_server/internal/domain/entity"

type IndexDTO struct {
	Index int `json:"index"`
}

func ConvertToIndexDTO(val *entity.Index) IndexDTO {
	return IndexDTO{
		Index: val.Index,
	}
}
