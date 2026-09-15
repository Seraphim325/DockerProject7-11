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
	client   *redis.Client
	prefix   string
	sentinel string
}

type redisValueModel struct {
	Key   string
	Value string
}

func NewValueRepository(client *redis.Client, prefix string, sentinel string) repository.ValueRepository {
	return &redisValueRepository{
		client:   client,
		prefix:   prefix,
		sentinel: sentinel,
	}
}

func (r *redisValueRepository) Save(ctx context.Context, value *entity.ValueRequest) (*entity.ValueResponse, error) {
	model := redisValueModel{
		Key:   fmt.Sprintf("%s%d", r.prefix, value.Value),
		Value: r.sentinel,
	}

	err := r.client.Set(ctx, model.Key, model.Value, 24*time.Hour).Err()

	if err != nil {
		return nil, err
	}

	return &entity.ValueResponse{
		Key:   model.Key,
		Value: model.Value,
	}, nil
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
