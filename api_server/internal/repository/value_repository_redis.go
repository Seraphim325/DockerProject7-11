package repository

import (
	"api_server/internal/domain/entity"
	"api_server/internal/domain/repository"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisValueRepository struct {
	client *redis.Client
}

func NewValueRepository(client *redis.Client) repository.ValueRepository {
	return &redisValueRepository{client: client}
}

func (r *redisValueRepository) Save(ctx context.Context, value *entity.ValueRepositoryRequest) error {
	err := r.client.Set(ctx, value.Key, value.Value, 24*time.Hour).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *redisValueRepository) GetAll(ctx context.Context) ([]entity.ValueResponse, error) {
	var values []entity.ValueResponse

	keys, err := r.getAllKeys(ctx)

	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		value, err := r.client.Get(ctx, key).Result()

		if err != nil {
			return nil, err
		}

		values = append(values, entity.ValueResponse{Key: key, Value: value})
	}

	return values, nil
}

func (r *redisValueRepository) getAllKeys(ctx context.Context) ([]string, error) {
	var keys []string
	var cursor uint64

	for {
		result, nextCursor, err := r.client.Scan(ctx, cursor, "*", 100).Result()

		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		keys = append(keys, result...)
		cursor = nextCursor

		if cursor == 0 {
			break
		}

	}

	return keys, nil
}
