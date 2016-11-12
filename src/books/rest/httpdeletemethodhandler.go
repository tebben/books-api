package rest

import (
	"github.com/tebben/books-api/src/books/models"
	"net/http"
)

// handleDeleteRequest
func handleDeleteRequest(w http.ResponseWriter, e *models.Endpoint, r *http.Request, h *func() error) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	handle := *h
	err := handle()
	if err != nil {
		sendError(w, []error{err})
		return
	}

	sendJSONResponse(w, http.StatusOK, nil)
}
