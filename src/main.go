package main

import (
	"log"

	"flag"
	"github.com/tebben/books-api/src/books/api"
	"github.com/tebben/books-api/src/books/database"
	"github.com/tebben/books-api/src/books/models"
	"github.com/tebben/books-api/src/http"
	"os"
	"strconv"
)

var (
	dbPath string
	ip     string
	port   int
)

func main() {
	initSettings()
	database := database.NewDatabase(dbPath)
	database.Start()
	stAPI := api.NewAPI(database)
	createAndStartServer(ip, port, &stAPI)
}

func initSettings() {
	dbFlag := flag.String("db", "books.db", "path to the database to use, can be overwritten by environment variable \"books_db\"")
	ipFlag := flag.String("h", "127.0.0.1", "local ip of the http server, can be overwritten by environment variable \"books_host\"")
	portFlag := flag.Int("p", 8080, "port of the http server, can be overwritten by environment variable \"books_port\"")
	flag.Parse()

	dbPath = *dbFlag
	ip = *ipFlag
	port = *portFlag

	envDB := os.Getenv("books_db")
	if envDB != "" {
		dbPath = envDB
	}

	envIP := os.Getenv("books_host")
	if envIP != "" {
		ip = envIP
	}

	envPort := os.Getenv("books_port")
	if envPort != "" {
		setPortFromString(envPort)
	}
}

func setPortFromString(newPort string) {
	i, e := strconv.Atoi(newPort)
	if e != nil {
		log.Fatal("Given port is not an int")
	}

	port = i
}

// createAndStartServer creates the books HTTPServer and starts it
func createAndStartServer(ip string, port int, api *models.API) {
	a := *api
	a.Start()

	booksServer := http.CreateServer(ip, port, api)
	booksServer.Start()
}
