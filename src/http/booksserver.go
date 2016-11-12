package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/tebben/books-api/src/books/models"
)

// Server interface for starting and stopping the HTTP server
type Server interface {
	Start()
	Stop()
}

// BooksServer is the type that contains all of the relevant information to set
// up the books HTTP Server
type BooksServer struct {
	host string      // Hostname for example "localhost" or "192.168.1.14"
	port int         // Port number where you want to run your http server on
	api  *models.API // books api to interact with from the HttpServer
}

// CreateServer initialises a new books HTTPServer based on the given parameters
func CreateServer(host string, port int, api *models.API) Server {
	return &BooksServer{
		host: host,
		port: port,
		api:  api,
	}
}

// Start command to start the books HTTPServer
func (s *BooksServer) Start() {
	log.Printf("Started books HTTP Server on %v:%v", s.host, s.port)
	router := CreateRouter(s.api)
	httpError := http.ListenAndServe(s.host+":"+strconv.Itoa(s.port), router)

	if httpError != nil {
		log.Fatal(httpError)
		return
	}
}

// Stop command to gracefully stop the books HTTP server, currently not supported
func (s *BooksServer) Stop() {

}
