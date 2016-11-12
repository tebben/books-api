package rest

import (
	"net/http"

	"github.com/tebben/books-api/src/books/models"
)

// handleGetRequest is the default function to handle incoming GET requests
func handleGetRequest(w http.ResponseWriter, e *models.Endpoint, r *http.Request, h *func() (interface{}, error)) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Run the handler func
	handler := *h
	data, err2 := handler()
	if err2 != nil {
		sendError(w, []error{err2})
		return
	}

	sendJSONResponse(w, http.StatusOK, data)
}
