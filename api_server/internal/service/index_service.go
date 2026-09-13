package service

import (
	"api_server/internal/domain"
	"api_server/internal/domain/entity"
	"api_server/internal/domain/repository"
	"context"
	"errors"
	"fmt"
)

type IndexService struct {
	repo repository.IndexRepository
}

func NewIndexService(repo repository.IndexRepository) *IndexService {
	return &IndexService{repo: repo}
}

func (s *IndexService) SaveIndex(ctx context.Context, val int) (*entity.Index, error) {

	if val > 50 {
		return nil, domain.ErrBigNumber
	}

	model := &entity.Index{Index: val}

	err := s.repo.Save(ctx, model)

	if err != nil {
		if errors.Is(err, domain.ErrDuplicate) {
			return nil, fmt.Errorf("number %d already exists: %w", val, err)
		}
		return nil, err
	}

	return model, nil
}

func (s *IndexService) GetAllIndexes(ctx context.Context) ([]entity.Index, error) {

	indexes, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to query database: %s", err)
	}

	return indexes, nil
}
