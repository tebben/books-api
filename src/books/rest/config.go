package rest

import "github.com/tebben/books-api/src/books/models"

func CreateEndPoints() []models.Endpoint {
	endpoints := []models.Endpoint{
		createBooksEndpoint(),
		createUsersEndpoint(),
	}

	return endpoints
}
