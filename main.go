package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string  `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author *Person `json:"author,omitempty"`
	Isbn   int64   `json:"isbn,omitempty"`
	Genre  string  `json:"genre,omitempty"`
}

type Person struct {
	Name   string `json:"name,omitempty"`
	Gneder string `json:"gneder,omitempty"`
	Age    int    `json:"age,omitempty"`
}

var Books []Book

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, book := range Books {

		if params["id"] == book.Id {
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)

}

func createBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000000))
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var bookToUpdate Book
	json.NewDecoder(r.Body).Decode(&bookToUpdate)

	for index, book := range Books {
		if book.Id == bookToUpdate.Id {
			Books = append(Books[:index], Books[index+1:]...)
			Books = append(Books, bookToUpdate)
			json.NewEncoder(w).Encode(bookToUpdate)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range Books {
		if book.Id == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not foud", http.StatusNotFound)
}

func main() {
	initialize()
	port := 8080
	handler := mux.NewRouter()

	handler.HandleFunc("/books", getAllBooks).Methods("GET")
	handler.HandleFunc("/book/{id}", getBook).Methods("GET")
	handler.HandleFunc("/book", createBook).Methods("POST")
	handler.HandleFunc("/book", updateBook).Methods("PUT")
	handler.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	fmt.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))

}

func initialize() {

	book1 := Book{
		Id:    "123",
		Title: "Can't hurt me",
		Isbn:  12302560,
		Genre: "Motivational",
		Author: &Person{
			Name:   "David Goggins",
			Gneder: "Male",
			Age:    45,
		},
	}

	book2 := Book{
		Id:    "777",
		Title: "Mrutunjay",
		Isbn:  2356,
		Genre: "Devotinal, Biopic",
		Author: &Person{
			Name:   "Shivaji Sawant",
			Gneder: "Male",
			Age:    75,
		},
	}

	Books = append(Books, book1, book2)

}
