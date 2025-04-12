package post

import (
	"encoding/json"
	"forum/service"
	"net/http"
	"strconv"
)

type ListPostHandler struct {
	service *service.PostService
}

func NewListPostHandler(service *service.PostService) *ListPostHandler {
	return &ListPostHandler{
		service: service,
	}
}

func (h *ListPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	limit := qp.Get("limit")
	offset := qp.Get("offset")
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Invalid limit", http.StatusBadRequest)
		return
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Invalid offset", http.StatusBadRequest)
		return
	}

	posts, err := h.service.List(r.Context(), limitInt, offsetInt)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
