package post

import (
	"forum/service"
	"net/http"
	"strconv"
)

type DeletePostHandler struct {
	service *service.PostService
}

func NewDeletePostHandler(service *service.PostService) *DeletePostHandler {
	return &DeletePostHandler{
		service: service,
	}
}

func (h *DeletePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	postId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(r.Context(), postId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
