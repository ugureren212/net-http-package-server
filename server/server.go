package main

import (
	"fmt"
	"net/http"
)

// TODO: add json encoding
type Book struct {
	name  string
	genre string
	price float32
}

func generateDB() []Book {
	db := make([]Book, 0)
	db = append(db, Book{name: "harry-potter", genre: "fantasy", price: 7.99})
	db = append(db, Book{name: "dune", genre: "adventure", price: 10.99})

	return db
}

func getAllBooksHandler(w http.ResponseWriter, req *http.Request, db []Book) {
	fmt.Fprintf(w, "\n")

	// TODO: fix book price/float bug. Returns value with too many decimal places
	for _, book := range db {
		msg := fmt.Sprintf("name: %s, genre: %s, price: %f", book.name, book.genre, book.price)
		fmt.Fprintf(w, msg, "\n")

	}

}

func main() {

	db := generateDB()

	fmt.Println(db)

	// TODO: Learn closure
	http.HandleFunc("/getAllBooks", func(w http.ResponseWriter, r *http.Request) {
		getAllBooksHandler(w, r, db)
	})

	http.ListenAndServe(":8090", nil)
}
