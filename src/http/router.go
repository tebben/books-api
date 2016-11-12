package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tebben/books-api/src/books/models"
)

// CreateRouter creates a new mux.Router and sets up all endpoints defined in the books api
func CreateRouter(api *models.API) *mux.Router {
	a := *api
	endpoints := *a.GetEndpoints()
	router := mux.NewRouter().StrictSlash(true)

	for _, endpoint := range endpoints {
		ep := endpoint
		for _, op := range ep.GetOperations() {
			operation := op
			method := fmt.Sprintf("%s", operation.OperationType)
			if operation.Handler == nil {
				continue
			}

			router.Methods(method).
				Path(operation.Path).
				HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					operation.Handler(w, r, &ep, api)
				})
		}
	}

	return router
}
