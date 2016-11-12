package rest

import (
	"github.com/tebben/books-api/src/books/models"
	"net/http"
)

func HandleGetUsers(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() (interface{}, error) { return a.GetUsers() }
	handleGetRequest(w, endpoint, r, &handle)
}

func HandlePostUser(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	u := &models.User{}
	handle := func() (interface{}, error) { return a.PostUser(u) }
	handlePostRequest(w, endpoint, r, u, &handle)
}

func HandlePatchUser(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	u := &models.User{}
	handle := func() (interface{}, error) { return a.PatchUser(getEntityID(r), u) }
	handlePatchRequest(w, endpoint, r, u, &handle)
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() error { return a.DeleteUser(getEntityID(r)) }
	handleDeleteRequest(w, endpoint, r, &handle)
}

func HandleGetUserByID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() (interface{}, error) { return a.GetUserByID(getEntityID(r)) }
	handleGetRequest(w, endpoint, r, &handle)
}

func HandleGetWishListByUserID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() (interface{}, error) { return a.GetWishList(getEntityID(r)) }
	handleGetRequest(w, endpoint, r, &handle)
}

func HandlePostWishListByUserID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	wi := &models.WishList{}
	handle := func() (interface{}, error) { return a.PostWishList(getEntityID(r), wi) }
	handlePostRequest(w, endpoint, r, wi, &handle)
}

func HandlePatchWishListByUserID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	wi := &models.WishList{}
	handle := func() (interface{}, error) { return a.PatchWishList(getEntityID(r), wi) }
	handlePatchRequest(w, endpoint, r, wi, &handle)
}

func HandleGetReadByUserID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	handle := func() (interface{}, error) { return a.GetRead(getEntityID(r)) }
	handleGetRequest(w, endpoint, r, &handle)
}

func HandlePostReadByUserID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	re := &models.Read{}
	handle := func() (interface{}, error) { return a.PostRead(getEntityID(r), re) }
	handlePostRequest(w, endpoint, r, re, &handle)
}

func HandlePatchReadByUserID(w http.ResponseWriter, r *http.Request, endpoint *models.Endpoint, api *models.API) {
	a := *api
	re := &models.Read{}
	handle := func() (interface{}, error) { return a.PatchRead(getEntityID(r), re) }
	handlePatchRequest(w, endpoint, r, re, &handle)
}
