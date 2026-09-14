package dto

import "api_server/internal/domain/entity"

type ValueRequestDTO struct {
	Value int `json:"value"`
}

func ConvertToValueRequestDTO(val *entity.ValueRequest) ValueRequestDTO {
	return ValueRequestDTO{
		Value: val.Value,
	}
}

type ValueResponseDTO struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ConvertToValueResponseDTO(val *entity.ValueResponse) ValueResponseDTO {
	return ValueResponseDTO{
		Key:   val.Key,
		Value: val.Value,
	}
}
