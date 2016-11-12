package models

import (
	"encoding/json"
	"net/http"

	"errors"
	booksErrors "github.com/tebben/books-api/src/errors"
)

// API describes all request and responses to fulfill the SensorThings API standard
type API interface {
	Start()

	GetEndpoints() *[]Endpoint

	GetBooks() ([]*Book, error)
	GetBookByID(id int) (*Book, error)
	PostBook(book *Book) (*Book, error)
	DeleteBook(id int) error
	PatchBook(id int, book *Book) (*Book, error)

	GetUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	PostUser(user *User) (*User, error)
	DeleteUser(id int) error
	PatchUser(id int, user *User) (*User, error)

	GetWishList(id int) (*WishList, error)
	PostWishList(id int, wishList *WishList) (*WishList, error)
	PatchWishList(id int, wishList *WishList) (*WishList, error)

	GetRead(id int) (*Read, error)
	PostRead(id int, read *Read) (*Read, error)
	PatchRead(id int, read *Read) (*Read, error)
}

// Database specifies the operations that the database provider needs to support
type Database interface {
	Start()
	CreateDB()

	GetBooks() ([]*Book, error)
	GetBookByID(id int) (*Book, error)
	PostBook(book *Book) (*Book, error)
	DeleteBook(id int) error
	PatchBook(id int, book *Book) (*Book, error)

	GetUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	PostUser(user *User) (*User, error)
	DeleteUser(id int) error
	PatchUser(id int, user *User) (*User, error)

	GetWishList(id int) (*WishList, error)
	PostWishList(id int, wishList *WishList) (*WishList, error)
	PatchWishList(id int, wishList *WishList) (*WishList, error)

	GetRead(id int) (*Read, error)
	PostRead(id int, read *Read) (*Read, error)
	PatchRead(id int, read *Read) (*Read, error)
}

// Endpoint defines the rest endpoint options
type Endpoint interface {
	GetName() string
	GetOperations() []EndpointOperation
}

// HTTPHandler func defines the format of the handler to process the incoming request
type HTTPHandler func(w http.ResponseWriter, r *http.Request, e *Endpoint, a *API)

// EndpointOperation contains the needed information to create an endpoint in the HTTP.Router
type EndpointOperation struct {
	OperationType HTTPOperation
	Path          string //relative path to the endpoint for example: /books
	Handler       HTTPHandler
}

// HTTPOperation describes the HTTP operation such as GET POST DELETE.
type HTTPOperation string

// HTTPOperation is a "enumeration" of the HTTP operations needed for all endpoints.
const (
	HTTPOperationGet    HTTPOperation = "GET"
	HTTPOperationPost   HTTPOperation = "POST"
	HTTPOperationPatch  HTTPOperation = "PATCH"
	HTTPOperationPut    HTTPOperation = "PUT"
	HTTPOperationDelete HTTPOperation = "DELETE"
)

// ArrayResponse is the default response format for sending content back
type ArrayResponse struct {
	Count    int          `json:"count,omitempty"`
	NextLink string       `json:"@iot.nextLink,omitempty"`
	Data     *interface{} `json:"value"`
}

// ErrorResponse is the default response format for sending errors back
type ErrorResponse struct {
	Error ErrorContent `json:"error"`
}

type Entity interface {
	ParseEntity(data []byte) error
}

type BaseEntity struct {
}

func (b *BaseEntity) ParseEntity(data []byte) error {
	return nil
}

// ErrorContent holds information on the error that occurred
type ErrorContent struct {
	StatusText string   `json:"status"`
	StatusCode int      `json:"code"`
	Messages   []string `json:"message"`
}

// Book describes a book object
type Book struct {
	BaseEntity
	ID       int    `json:"id"`
	Title    string `json:"title,omitempty"`
	Category string `json:"category,omitempty"`
	Pages    int    `json:"pages,omitempty"`
	ISBN     string `json:"ISBN,omitempty"`
}

// ParseEntity tries to parse the given json byte array into the current entity
func (b *Book) ParseEntity(data []byte) error {
	book := &b
	err := json.Unmarshal(data, book)
	if err != nil {
		return booksErrors.NewBadRequestError(errors.New("Unable to parse Datastream"))
	}

	return nil
}

// User describes a user object
type User struct {
	BaseEntity
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

// ParseEntity tries to parse the given json byte array into the current entity
func (u *User) ParseEntity(data []byte) error {
	user := &u
	err := json.Unmarshal(data, user)
	if err != nil {
		return booksErrors.NewBadRequestError(errors.New("Unable to parse user"))
	}

	return nil
}

// WishList describes a WishList object
type WishList struct {
	BaseEntity
	Books []int `json:"wishList"`
}

// ParseEntity tries to parse the given json byte array into the current entity
func (w *WishList) ParseEntity(data []byte) error {
	wish := &w
	err := json.Unmarshal(data, wish)
	if err != nil {
		return booksErrors.NewBadRequestError(errors.New("Unable to parse wishlist"))
	}

	return nil
}

// Read describes a Read object
type Read struct {
	BaseEntity
	Books []int `json:"read"`
}

// ParseEntity tries to parse the given json byte array into the current entity
func (r *Read) ParseEntity(data []byte) error {
	read := &r
	err := json.Unmarshal(data, read)
	if err != nil {
		return booksErrors.NewBadRequestError(errors.New("Unable to parse read list"))
	}

	return nil
}
