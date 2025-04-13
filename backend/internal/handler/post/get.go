package post

import (
	"encoding/json"
	"forum/internal/service"
	"net/http"
	"strconv"
)

type GetPostHandler struct {
	service *service.PostService
}

func NewGetPostHandler(service *service.PostService) *GetPostHandler {
	return &GetPostHandler{
		service: service,
	}
}

func (h *GetPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.service.Get(r.Context(), postId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
