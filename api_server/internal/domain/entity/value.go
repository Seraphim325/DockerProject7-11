package entity

import "github.com/redis/go-redis/v9"

type ValueRequest struct {
	Value int
}

type ValueResponse struct {
	Key   string
	Value string
}

type ValueRepositoryRequest struct {
	Key   string
	Value redis.Error
}
