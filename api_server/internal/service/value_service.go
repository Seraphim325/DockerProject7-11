package service

import (
	"api_server/internal/domain"
	"api_server/internal/domain/entity"
	"api_server/internal/domain/repository"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type ValueService struct {
	repo   repository.ValueRepository
	prefix string
}

func NewValueService(repo repository.ValueRepository, prefix string) *ValueService {
	return &ValueService{
		repo:   repo,
		prefix: prefix,
	}
}

func (s *ValueService) SaveValue(ctx context.Context, val int) (*entity.ValueResponse, error) {

	if val > 50 {
		return nil, domain.ErrBigNumber
	}

	model := &entity.ValueRepositoryRequest{Key: fmt.Sprintf("%s%d", s.prefix, val), Value: redis.Nil}

	err := s.repo.Save(ctx, model)

	if err != nil {
		return nil, err
	}

	response := &entity.ValueResponse{Key: model.Key}

	return response, nil
}

func (s *ValueService) GetAllValues(ctx context.Context) ([]entity.ValueResponse, error) {

	values, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return values, nil
}
