package repository

import (
	"api_server/internal/domain"
	"api_server/internal/domain/entity"
	"api_server/internal/domain/repository"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type GormIndexModel struct {
	ID    int64 `gorm:"primaryKey"`
	Index int   `gorm:"uniqueIndex"`
}

func (GormIndexModel) TableName() string {
	return "indexes"
}

type GormIndexRepository struct {
	db *gorm.DB
}

func NewGormIndexRepository(db *gorm.DB) repository.IndexRepository {
	return &GormIndexRepository{db: db}
}

func (r *GormIndexRepository) Save(ctx context.Context, index *entity.Index) error {
	model := GormIndexModel{
		Index: index.Index,
	}

	if err := r.db.WithContext(ctx).Save(&model).Error; isDuplicateError(err) {
		return domain.ErrDuplicate
	}

	return nil
}

func (r *GormIndexRepository) GetAll(ctx context.Context) ([]entity.Index, error) {
	var models []GormIndexModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	indexes := make([]entity.Index, 0, len(models))

	for _, model := range models {
		indexes = append(indexes, entity.Index{
			ID:    model.ID,
			Index: model.Index,
		})
	}

	return indexes, nil
}

func isDuplicateError(err error) bool {
	if err == nil {
		return false
	}

	pgErr, ok := err.(*pgconn.PgError)

	if ok {
		return pgErr.Code == "23505"
	} else {
		return false
	}
}
