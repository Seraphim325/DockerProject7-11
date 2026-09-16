package service

import (
	"api_server/internal/domain"
	"api_server/internal/domain/entity"
	"api_server/internal/domain/repository"
	"context"
)

type ValueService struct {
	repo repository.ValueRepository
}

func NewValueService(repo repository.ValueRepository) *ValueService {
	return &ValueService{
		repo: repo,
	}
}

func (s *ValueService) SaveValue(ctx context.Context, val int) (*entity.ValueResponse, error) {

	if val > domain.ThresholdValue {
		return nil, domain.ErrBigNumber
	}

	model := &entity.ValueRequest{Value: val}

	resp, err := s.repo.Save(ctx, model)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ValueService) GetAllValues(ctx context.Context) ([]entity.ValueResponse, error) {

	values, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return values, nil
}
