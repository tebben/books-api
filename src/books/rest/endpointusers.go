package rest

import "github.com/tebben/books-api/src/books/models"

func createUsersEndpoint() *Endpoint {
	return &Endpoint{
		Name: "Users",
		Operations: []models.EndpointOperation{
			{models.HTTPOperationGet, "/Users", HandleGetUsers},
			{models.HTTPOperationPost, "/Users", HandlePostUser},

			{models.HTTPOperationGet, "/Users/{id}", HandleGetUserByID},
			{models.HTTPOperationPatch, "/Users/{id}", HandlePatchUser},
			{models.HTTPOperationPut, "/Users/{id}", HandlePatchUser},
			{models.HTTPOperationDelete, "/Users/{id}", HandleDeleteUser},

			{models.HTTPOperationGet, "/Users/{id}/WishList", HandleGetWishListByUserID},
			{models.HTTPOperationPost, "/Users/{id}/WishList", HandlePostWishListByUserID},
			{models.HTTPOperationPatch, "/Users/{id}/WishList", HandlePatchWishListByUserID},
			{models.HTTPOperationPut, "/Users/{id}/WishList", HandlePatchWishListByUserID},

			{models.HTTPOperationGet, "/Users/{id}/Read", HandleGetReadByUserID},
			{models.HTTPOperationPost, "/Users/{id}/Read", HandlePostReadByUserID},
			{models.HTTPOperationPatch, "/Users/{id}/Read", HandlePatchReadByUserID},
			{models.HTTPOperationPut, "/Users/{id}/Read", HandlePatchReadByUserID},
		},
	}
}
