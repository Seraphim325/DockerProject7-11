package dto

import "api_server/internal/domain/entity"

type ValueRequestDTO struct {
	Value int
}

func ConvertToValueRequestDTO(val *entity.ValueRequest) ValueRequestDTO {
	return ValueRequestDTO{
		Value: val.Value,
	}
}

type ValueResponseDTO struct {
	Key   string
	Value string
}

func ConvertToValueResponseDTO(val *entity.ValueResponse) ValueResponseDTO {
	return ValueResponseDTO{
		Key:   val.Key,
		Value: val.Value,
	}
}
