package index

import (
	"api_server/internal/service"
	"api_server/internal/transport/http/handler/index/dto"
	handler "api_server/internal/transport/http/handler/shared"
	"encoding/json"
	"net/http"
)

type IndexHandler struct {
	service *service.IndexService
}

func NewIndexService(service *service.IndexService) *IndexHandler {
	return &IndexHandler{service: service}
}

func (h *IndexHandler) SaveIndex(w http.ResponseWriter, r *http.Request) {
	var req dto.IndexDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.Error(w, err)
		return
	}

	index, err := h.service.SaveIndex(r.Context(), req.Index)

	if err != nil {
		handler.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.ConvertToIndexDTO(index))
}

func (h *IndexHandler) GetAllIndexes(w http.ResponseWriter, r *http.Request) {
	indexes, err := h.service.GetAllIndexes(r.Context())

	if err != nil {
		handler.Error(w, err)
		return
	}

	resp := make([]dto.IndexDTO, 0, len(indexes))

	for _, index := range indexes {
		resp = append(resp, dto.ConvertToIndexDTO(&index))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
