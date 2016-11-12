package api

import (
	"github.com/tebben/books-api/src/books/models"
	"github.com/tebben/books-api/src/books/rest"
)

// BookAPI is the default implementation of books, API needs a database
// provider, endpoint information to setup te needed services
type BooksAPI struct {
	db        models.Database
	endPoints []models.Endpoint
}

// NewAPI Initialise a new books API
func NewAPI(database models.Database) models.API {
	return &BooksAPI{
		db: database,
	}
}

// Start is used to set the initial state of the api
func (a *BooksAPI) Start() {
}

// GetEndpoints returns all configured endpoints for the HTTP server
func (a *BooksAPI) GetEndpoints() *[]models.Endpoint {
	if a.endPoints == nil {
		a.endPoints = rest.CreateEndPoints()
	}

	return &a.endPoints
}
