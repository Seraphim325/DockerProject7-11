package main

import (
	"api_server/internal/config"
	"api_server/internal/database"
	"api_server/internal/redis"
	"api_server/internal/repository"
	"api_server/internal/service"
	t_http "api_server/internal/transport/http"
	"api_server/internal/transport/http/handler"
	"context"
	"fmt"
	"log"

	"net/http"
)

func main() {
	conf, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load config: %s\n", err)
	}

	ctx := context.Background()

	client, err := redis.NewClient(ctx, conf.RedisConfig)

	if err != nil {
		log.Fatalf("Failed to connect to redis db: %s\n", err)
	}

	db, err := database.NewDB(conf.DatabaseConfig)

	if err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err)
	}

	if err := db.AutoMigrate(
		&repository.GormIndexModel{},
	); err != nil {
		log.Fatalf("Failed to migrate: %s\n", err)
	}

	indexRepository := repository.NewGormIndexRepository(db)
	valueRepository := repository.NewValueRepository(client, conf.RedisConfig.Pattern, conf.RedisConfig.Sentinel)

	indexService := service.NewIndexService(indexRepository)
	valueService := service.NewValueService(valueRepository)

	handler := handler.NewHandler(indexService, valueService)

	router := t_http.NewRouter(handler)

	log.Printf("Running server on port %s \n", conf.AppConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.AppConfig.Port), router)

}
