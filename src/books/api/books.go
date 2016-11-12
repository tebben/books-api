package api

import (
	"github.com/tebben/books-api/src/books/models"
)

// GetBooks retrieves requests all books from the database
func (a *BooksAPI) GetBooks() ([]*models.Book, error) {
	return a.db.GetBooks()
}

// GetBookByID requests a book from the database by given id
func (a *BooksAPI) GetBookByID(id int) (*models.Book, error) {
	return a.db.GetBookByID(id)
}

func (a *BooksAPI) PostBook(book *models.Book) (*models.Book, error) {
	return a.db.PostBook(book)
}

func (a *BooksAPI) DeleteBook(id int) error {
	return a.db.DeleteBook(id)
}

func (a *BooksAPI) PatchBook(id int, book *models.Book) (*models.Book, error) {
	return a.db.PatchBook(id, book)
}
