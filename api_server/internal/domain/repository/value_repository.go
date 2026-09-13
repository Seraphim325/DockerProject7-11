package repository

import (
	"api_server/internal/domain/entity"
	"context"
)

type ValueRepository interface {
	Save(ctx context.Context, value *entity.ValueRepositoryRequest) error
	GetAll(ctx context.Context) ([]entity.ValueResponse, error)
}
