package database

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/tebben/books-api/src/books/models"
	"log"
)

var (
	UsersBucket    = "users"
	BooksBucket    = "books"
	WishListBucket = "wishlist"
	ReadBucket     = "read"
)

type BookStore struct {
	DB   *bolt.DB
	Path string
}

func NewDatabase(path string) models.Database {
	return &BookStore{
		Path: path,
	}
}

// CreateDB creates the database if not exist
func (bdb *BookStore) CreateDB() {
	db, err := bolt.Open(bdb.Path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(UsersBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte(BooksBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte(WishListBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte(ReadBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal("Unable to create database buckets")
	}

	//defer db.Close()

	bdb.DB = db
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (bdb *BookStore) Start() {
	bdb.CreateDB()
}

func (bdb *BookStore) Stop() {

}
