package rest

import "github.com/tebben/books-api/src/books/models"

func createBooksEndpoint() *Endpoint {
	return &Endpoint{
		Name: "Books",
		Operations: []models.EndpointOperation{
			{models.HTTPOperationGet, "/Books", HandleGetBooks},
			{models.HTTPOperationPost, "/Books", HandlePostBook},

			{models.HTTPOperationGet, "/Books/{id}", HandleGetBookByID},
			{models.HTTPOperationPatch, "/Books/{id}", HandlePatchBook},
			{models.HTTPOperationPut, "/Books/{id}", HandlePatchBook},
			{models.HTTPOperationDelete, "/Books/{id}", HandleDeleteBook},
		},
	}
}
