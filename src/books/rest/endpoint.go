package rest

import "github.com/tebben/books-api/src/books/models"

// Endpoint bla
type Endpoint struct {
	Name       string                     `json:"name"` // Name of the endpoint
	Operations []models.EndpointOperation `json:"-"`
}

// GetName returns the endpoint name
func (e *Endpoint) GetName() string {
	return e.Name
}

// GetOperations returns all operations for this endpoint such as GET, POST
func (e *Endpoint) GetOperations() []models.EndpointOperation {
	return e.Operations
}
