package repository

import (
	"api_server/internal/domain/entity"
	"context"
)

type IndexRepository interface {
	Save(ctx context.Context, index *entity.Index) error
	GetAll(ctx context.Context) ([]entity.Index, error)
}
