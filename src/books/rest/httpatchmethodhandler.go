package rest

import (
	"github.com/tebben/books-api/src/books/models"
	"io/ioutil"
	"net/http"
)

// handlePatchRequest todo: currently almost same as handlePostRequest, merge if it stays like this
func handlePatchRequest(w http.ResponseWriter, e *models.Endpoint, r *http.Request, entity models.Entity, h *func() (interface{}, error)) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if !checkContentType(w, r) {
		return
	}

	byteData, _ := ioutil.ReadAll(r.Body)
	err := entity.ParseEntity(byteData)
	if err != nil {
		sendError(w, []error{err})
		return
	}

	handle := *h
	data, err2 := handle()
	if err2 != nil {
		sendError(w, []error{err2})
		return
	}

	sendJSONResponse(w, http.StatusOK, data)
}
