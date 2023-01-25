package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"github.com/gorilla/mux" // used to get the params from the route
	"goCrudDemo/models"      // models package where User schema is defined
	"net/http"               // used to access the request and response object of the api
	"strconv"                // package used to covert string into int type
	"time"

	_ "github.com/lib/pq" // postgres golang driver
)

//type JsonResponse struct {
//	Type    string `json:"type"`
//	Message string `json:"message"`
//}
type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func setupDB() *sql.DB {
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", "dbname=books_database user=postgres password=Unnathi07$ host=localhost sslmode=disable") //sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	checkErr(err)
	return db
}

//Get all Books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi, I'm error")
	db := setupDB()

	fmt.Println("Getting book...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM books")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var books []models.Book

	// Foreach movie
	for rows.Next() {
		var id int
		var name string
		var author string
		var pages int
		var publication_date time.Time

		err = rows.Scan(&id, &name, &author, &pages, &publication_date)

		// check errors
		checkErr(err)
		currentBook := models.Book{Id: id, Name: name, Author: author, Pages: pages, DOP: publication_date}
		books = append(books, currentBook)
	}
	//msg := fmt.Sprintf(books)
	//res := response{
	//	ID:      int64(id),
	//	Message: msg,
	//}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	json.NewEncoder(w).Encode(books)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// create an empty book
	db := setupDB()
	var book models.Book

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&book)
	checkErr(err)
	var bookID int
	err = db.QueryRow(`INSERT INTO books(name, author, pages, publication_date) VALUES($1, $2, $3, $4) RETURNING id`, book.Name, book.Author, book.Pages, book.DOP).Scan(&bookID)

	//if err != nil {
	//	return 0, err
	//}
	checkErr((err))
	fmt.Printf("Last inserted ID: %v\n", bookID)
	res := response{
		ID:      int(bookID),
		Message: "added book",
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)
	var book models.Book
	// decode the json request to book
	err = json.NewDecoder(r.Body).Decode(&book)
	checkErr(err)
	//var updatedIds int
	res, err := db.Exec(`UPDATE books set name=$1, author=$2, pages=$3, publication_date=$4 where id=$5 RETURNING id`, book.Name, book.Author, book.Pages, book.DOP, id) //.Scan(&updatedIds)
	//if err != nil {
	//	return 0, err
	//}

	rowsUpdated, err := res.RowsAffected()
	fmt.Println(rowsUpdated)
	//if err != nil {
	//	return 0, err
	//}
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", rowsUpdated)
	checkErr(err)
	//res = JsonResponse{Type: "success", Message: "updated"}
	res1 := response{
		ID:      int(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res1)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)
	//var del int
	res, err := db.Exec(`delete from books where id = $1`, id) //.Scanf(&del)
	checkErr(err)
	deletedRows, err := res.RowsAffected()
	fmt.Println(deletedRows)
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the response message
	//res = JsonResponse{Type: "success", Message: "Deleted"}
	res1 := response{
		ID:      int(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res1)
}

//---------------------------------------------------------------------------------------------
//GetBook

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi, I'm error")
	params := mux.Vars(r)

	// convert the id type from string to int
	bookId, err := strconv.Atoi(params["id"])
	fmt.Println(bookId)
	checkErr(err)
	db := setupDB()
	// create an empty user of type models.User
	//var book models.Book

	// decode the json request to user
	//err = json.NewDecoder(r.Body).Decode(&book)

	checkErr(err)
	res := models.Book{}

	var id int
	var name string
	var author string
	var pages int
	var publicationDate time.Time

	err = db.QueryRow(`SELECT id, name, author, pages, publication_date FROM books where id = $1`, bookId).Scan(&id, &name, &author, &pages, &publicationDate)
	if err == nil {
		res = models.Book{Id: id, Name: name, Author: author, Pages: pages, DOP: publicationDate}
	}
	fmt.Println(res)
	// send the response
	json.NewEncoder(w).Encode(res)
}
