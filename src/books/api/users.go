package api

import "github.com/tebben/books-api/src/books/models"

// GetBooks retrieves requests all books from the database
func (a *BooksAPI) GetUsers() ([]*models.User, error) {
	return a.db.GetUsers()
}

// GetBookByID requests a book from the database by given id
func (a *BooksAPI) GetUserByID(id int) (*models.User, error) {
	return a.db.GetUserByID(id)
}

func (a *BooksAPI) PostUser(user *models.User) (*models.User, error) {
	return a.db.PostUser(user)
}

func (a *BooksAPI) DeleteUser(id int) error {
	return a.db.DeleteUser(id)
}

func (a *BooksAPI) PatchUser(id int, user *models.User) (*models.User, error) {
	return a.db.PatchUser(id, user)
}

func (a *BooksAPI) GetWishList(id int) (*models.WishList, error) {
	return a.db.GetWishList(id)
}

func (a *BooksAPI) PostWishList(id int, wl *models.WishList) (*models.WishList, error) {
	return a.db.PostWishList(id, wl)
}

func (a *BooksAPI) PatchWishList(id int, wl *models.WishList) (*models.WishList, error) {
	return a.db.PatchWishList(id, wl)
}

func (a *BooksAPI) GetRead(id int) (*models.Read, error) {
	return a.db.GetRead(id)
}

func (a *BooksAPI) PostRead(id int, r *models.Read) (*models.Read, error) {
	return a.db.PostRead(id, r)
}

func (a *BooksAPI) PatchRead(id int, r *models.Read) (*models.Read, error) {
	return a.db.PatchRead(id, r)
}
