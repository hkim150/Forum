package post

import (
	"encoding/json"
	"forum/internal/service"
	"net/http"
	"strconv"
)

type UpdatePostHandler struct {
	service *service.PostService
}

func NewUpdatePostHandler(service *service.PostService) *UpdatePostHandler {
	return &UpdatePostHandler{
		service: service,
	}
}

type UpdatePostParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *UpdatePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	params := UpdatePostParams{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.service.Update(r.Context(), postId, params.Title, params.Content)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
