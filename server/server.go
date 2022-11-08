// TODO: Improvements -> https://www.youtube.com/watch?v=nZaHOkiQCzw&t=497s&ab_channel=Fabrzy
// TODO: Gorilla mux -> https://www.youtube.com/watch?v=5BIylxkudaE&ab_channel=GoWebExamples

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// TODO: add json encoding
type Book struct {
	name  string
	genre string
	price float64
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
	nameParam := req.URL.Query().Get("name")

	for _, book := range db {
		if book.name == nameParam {
			bookPriceRounded := fmt.Sprintf("%.2f", book.price)
			msg := fmt.Sprintf("name: %s | genre: %s | price: %s \n", book.name, book.genre, bookPriceRounded)
			fmt.Fprintf(w, msg)
			return
		}
	}

	msg := fmt.Sprint(w, "cannot find book of name ", nameParam)
	fmt.Fprintf(w, msg)
}

func setBookHandler(w http.ResponseWriter, req *http.Request, db *[]Book) {
	//check if request type is post
	//unmarshal json data
	//append to data base
	nameParam := req.URL.Query().Get("name")
	genreParam := req.URL.Query().Get("genre")
	priceParam := req.URL.Query().Get("price")

	fmt.Println(nameParam, genreParam, priceParam)
	
	floatPriceParam, err := strconv.ParseFloat(priceParam, 64)
	if err != nil {
		// TODO: add proper error handling
		fmt.Println("Could not convert book price (string) into float")
	}

	*db = append(*db, Book{name: nameParam, genre: genreParam, price: floatPriceParam})
}

// TODO: add proper error response
func main() {

	db := generateDB()

	fmt.Println(db)

	// TODO: Learn closure
	http.HandleFunc("/getAllBooks", func(w http.ResponseWriter, r *http.Request) {
		getAllBooksHandler(w, r, db)
	})

	http.HandleFunc("/getBook", func(w http.ResponseWriter, r *http.Request) {
		getBookHandler(w, r, db)
	})

	http.HandleFunc("/setBook", func(w http.ResponseWriter, r *http.Request) {
		setBookHandler(w, r, &db)
		fmt.Println(db)
	})

	http.ListenAndServe(":8090", nil)
}

// curl http://localhost:8090/getBook?name=harry-potter