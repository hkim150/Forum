package post

import (
	"encoding/json"
	"forum/service"
	"net/http"
)

type CreatePostHandler struct {
	service *service.PostService
}

func NewCreatePostHandler(service *service.PostService) *CreatePostHandler {
	return &CreatePostHandler{
		service: service,
	}
}

type CreatePostParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *CreatePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := CreatePostParams{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(r.Context(), params.Title, params.Content)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
