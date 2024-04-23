package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}
type User struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

var books []Book = []Book{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", addBook).Methods("POST")
	r.HandleFunc("/books", getAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", patchBook).Methods("PATCH")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":5055", r)
}
func addBook(w http.ResponseWriter, r *http.Request) {
	var newbook Book
	json.NewDecoder(r.Body).Decode(&newbook)
	books = append(books, newbook)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	idparm := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparm)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("no book found with specified id"))
		return
	}
	book := books[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	idparm := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparm)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("no book found with specified id"))
		return
	}
	var updatebook Book
	json.NewDecoder(r.Body).Decode(&updatebook)
	books[id] = updatebook
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatebook)
}

func patchBook(w http.ResponseWriter, r *http.Request) {
	idparm := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparm)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("no book found with specified id"))
		return
	}
	book := &books[id]
	json.NewDecoder(r.Body).Decode(book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	idparm := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparm)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("no book found with specified id"))
		return
	}
	books = append(books[:id], books[id+1:]...)
	w.WriteHeader(200)
}
