// TODO: Improvements -> https://www.youtube.com/watch?v=nZaHOkiQCzw&t=497s&ab_channel=Fabrzy
// TODO: Gorilla mux -> https://www.youtube.com/watch?v=5BIylxkudaE&ab_channel=GoWebExamples

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
	for _, book := range db {
		bookPriceRounded := fmt.Sprintf("%.2f", book.price)
		msg := fmt.Sprintf("name: %s | genre: %s | price: %s \n", book.name, book.genre, bookPriceRounded)
		fmt.Fprintf(w, msg)
	}
}

func getBookHandler(w http.ResponseWriter, req *http.Request, db []Book) {
	param1 := req.URL.Query().Get("name")

	for _, book := range db {
		if book.name == param1 {
			bookPriceRounded := fmt.Sprintf("%.2f", book.price)
			msg := fmt.Sprintf("name: %s | genre: %s | price: %s \n", book.name, book.genre, bookPriceRounded)
			fmt.Fprintf(w, msg)
		} else {
			msg := fmt.Sprint(w, "cannot find book of name ", param1)
			fmt.Fprintf(w, "cannot find book of name ", msg)
		}
	}
}

func main() {

	// TODO: add proper response
	db := generateDB()

	fmt.Println(db)

	// TODO: Learn closure
	http.HandleFunc("/getAllBooks", func(w http.ResponseWriter, r *http.Request) {
		getAllBooksHandler(w, r, db)
	})

	http.HandleFunc("/getBook", func(w http.ResponseWriter, r *http.Request) {
		getBookHandler(w, r, db)
	})

	http.ListenAndServe(":8090", nil)
}
