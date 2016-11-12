package database

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/tebben/books-api/src/books/models"
)

func (bdb *BookStore) GetBooks() ([]*models.Book, error) {
	books := []*models.Book{}

	err := bdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BooksBucket))
		b.ForEach(func(k, v []byte) error {
			book := models.Book{}
			json.Unmarshal(v, &book)
			books = append(books, &book)

			return nil
		})
		return nil
	})

	return books, err
}

func (bdb *BookStore) GetBookByID(id int) (*models.Book, error) {
	book := models.Book{}

	err := bdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BooksBucket))
		v := b.Get(itob(id))
		json.Unmarshal(v, &book)
		return nil
	})

	return &book, err
}

func (bdb *BookStore) PostBook(book *models.Book) (*models.Book, error) {
	e := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BooksBucket))
		id, _ := b.NextSequence()
		book.ID = int(id)

		buf, err := json.Marshal(book)
		if err != nil {
			return err
		}

		err = b.Put(itob(book.ID), buf)
		return err
	})

	return book, e
}

func (bdb *BookStore) DeleteBook(id int) error {
	err := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BooksBucket))

		if err := b.Delete(itob(id)); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (bdb *BookStore) PatchBook(id int, book *models.Book) (*models.Book, error) {
	e := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BooksBucket))
		book.ID = id
		buf, err := json.Marshal(book)
		if err != nil {
			return err
		}

		err = b.Put(itob(book.ID), buf)
		return err
	})

	return book, e
}
