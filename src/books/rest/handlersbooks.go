package rest

import (
	"net/http"

	"github.com/tebben/books-api/src/books/models"
)

func HandleGetBooks(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() (interface{}, error) { return a.GetBooks() }
	handleGetRequest(w, endpoint, r, &handle)
}

func HandlePostBook(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	b := &models.Book{}
	handle := func() (interface{}, error) { return a.PostBook(b) }
	handlePostRequest(w, endpoint, r, b, &handle)
}

func HandlePatchBook(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	book := &models.Book{}
	handle := func() (interface{}, error) { return a.PatchBook(getEntityID(r), book) }
	handlePatchRequest(w, endpoint, r, book, &handle)
}

func HandleDeleteBook(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() error { return a.DeleteBook(getEntityID(r)) }
	handleDeleteRequest(w, endpoint, r, &handle)
}

func HandleGetBookByID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() (interface{}, error) { return a.GetBookByID(getEntityID(r)) }
	handleGetRequest(w, endpoint, r, &handle)
}
