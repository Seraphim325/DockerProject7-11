package handler

import (
	"api_server/internal/service"
	"api_server/internal/transport/http/handler/index"
	"api_server/internal/transport/http/handler/value"
)

type Handler struct {
	IndexHandler *index.IndexHandler
	ValueHandler *value.ValueHandler
}

func NewHandler(indexService *service.IndexService, valueService *service.ValueService) *Handler {
	return &Handler{
		IndexHandler: index.NewIndexService(indexService),
		ValueHandler: value.NewValueHandler(valueService),
	}
}
