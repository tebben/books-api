package rest

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	gostErrors "github.com/tebben/books-api/src/errors"
)

// getEntityID retrieves the id from the request
func getEntityID(r *http.Request) int {
	vars := mux.Vars(r)
	value := vars["id"]
	i, _ := strconv.Atoi(value)

	return i
}

func checkContentType(w http.ResponseWriter, r *http.Request) bool {
	// maybe needs to add case-insentive check?
	if len(r.Header.Get("Content-Type")) > 0 {
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			sendError(w, []error{gostErrors.NewBadRequestError(errors.New("Missing or wrong Content-Type, accepting: application/json"))})
			return false
		}
	}

	return true
}
