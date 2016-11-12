package database

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/tebben/books-api/src/books/models"
)

func (bdb *BookStore) GetUsers() ([]*models.User, error) {
	users := []*models.User{}

	err := bdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))
		b.ForEach(func(k, v []byte) error {
			user := models.User{}
			json.Unmarshal(v, &user)
			users = append(users, &user)

			return nil
		})
		return nil
	})

	return users, err
}

func (bdb *BookStore) GetUserByID(id int) (*models.User, error) {
	user := models.User{}

	err := bdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))
		v := b.Get(itob(id))
		json.Unmarshal(v, &user)
		return nil
	})

	return &user, err
}

func (bdb *BookStore) PostUser(user *models.User) (*models.User, error) {
	e := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))
		id, _ := b.NextSequence()
		user.ID = int(id)

		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}

		err = b.Put(itob(user.ID), buf)
		return err
	})

	return user, e
}

func (bdb *BookStore) DeleteUser(id int) error {
	err := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))

		if err := b.Delete(itob(id)); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (bdb *BookStore) PatchUser(id int, user *models.User) (*models.User, error) {
	e := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))
		user.ID = id
		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}

		err = b.Put(itob(user.ID), buf)
		return err
	})

	return user, e
}

func (bdb *BookStore) GetWishList(id int) (*models.WishList, error) {
	wl := models.WishList{}

	err := bdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(WishListBucket))
		v := b.Get(itob(id))
		json.Unmarshal(v, &wl)
		return nil
	})

	return &wl, err
}

func (bdb *BookStore) PostWishList(id int, wl *models.WishList) (*models.WishList, error) {
	e := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(WishListBucket))
		buf, err := json.Marshal(wl)
		if err != nil {
			return err
		}

		err = b.Put(itob(id), buf)
		return err
	})

	return wl, e
}

func (bdb *BookStore) PatchWishList(id int, wl *models.WishList) (*models.WishList, error) {
	return bdb.PostWishList(id, wl)
}

func (bdb *BookStore) GetRead(id int) (*models.Read, error) {
	r := models.Read{}

	err := bdb.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ReadBucket))
		v := b.Get(itob(id))
		json.Unmarshal(v, &r)
		return nil
	})

	return &r, err
}

func (bdb *BookStore) PostRead(id int, r *models.Read) (*models.Read, error) {
	e := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ReadBucket))
		buf, err := json.Marshal(r)
		if err != nil {
			return err
		}

		err = b.Put(itob(id), buf)
		return err
	})

	return r, e
}

func (bdb *BookStore) PatchRead(id int, r *models.Read) (*models.Read, error) {
	return bdb.PostRead(id, r)
}
