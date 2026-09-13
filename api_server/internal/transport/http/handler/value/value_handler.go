package value

import (
	"api_server/internal/service"
	handler "api_server/internal/transport/http/handler/shared"
	"api_server/internal/transport/http/handler/value/dto"
	"encoding/json"
	"net/http"
)

type ValueHandler struct {
	service *service.ValueService
}

func NewValueHandler(service *service.ValueService) *ValueHandler {
	return &ValueHandler{service: service}
}

func (h *ValueHandler) SaveValue(w http.ResponseWriter, r *http.Request) {
	var req dto.ValueRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.Error(w, err)
		return
	}

	value, err := h.service.SaveValue(r.Context(), req.Value)

	if err != nil {
		handler.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.ConvertToValueResponseDTO(value))
}

func (h *ValueHandler) GetAllValues(w http.ResponseWriter, r *http.Request) {
	values, err := h.service.GetAllValues(r.Context())

	if err != nil {
		handler.Error(w, err)
		return
	}

	resp := make([]dto.ValueResponseDTO, 0, len(values))

	for _, value := range values {
		resp = append(resp, dto.ConvertToValueResponseDTO(&value))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
