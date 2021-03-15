package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Models / Structs
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:author`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

// Handlers
func getBooks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, book := range books {
		if book.ID == params["id"] {
			var book Book
			_ = json.NewDecoder(req.Body).Decode(&book)
			book.ID = params["id"]
			books[index] = book
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(map[string]string{"detail": "No post found with specified id"})
}

func deleteBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Init Dummy Data
	books = append(books, Book{ID: "1", ISBN: "45454353", Title: "Book One", Author: &Author{
		Firstname: "John", Lastname: "Smith"}})
	books = append(books, Book{ID: "2", ISBN: "34534534", Title: "Book Two", Author: &Author{
		Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "3", ISBN: "34434534", Title: "Book Three", Author: &Author{
		Firstname: "John", Lastname: "Wick"}})
	books = append(books, Book{ID: "4", ISBN: "35367534", Title: "Book Four", Author: &Author{
		Firstname: "John", Lastname: "Carpenter"}})

	// Router Handlers / Endpoints
	router.HandleFunc("/api/book/v1/books", getBooks).Methods("GET")
	router.HandleFunc("/api/book/v1/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/book/v1/books", createBook).Methods("POST")
	router.HandleFunc("/api/book/v1/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/book/v1/books/{id}", deleteBook).Methods("DELETE")

	// Start Server
	log.Fatal(http.ListenAndServe(":8001", router))
}
