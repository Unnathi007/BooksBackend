package routers

import (
	"goCrudDemo/middleware"

	"fmt"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Server at 8080")
	router := mux.NewRouter()
	fmt.Println(router)
	router.HandleFunc("/api/book/{id}", middleware.GetBook).Methods("GET", "OPTIONS") //.Schemes("http") //.Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books", middleware.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/new-book", middleware.CreateBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/update-book/{id}", middleware.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete-book/{id}", middleware.DeleteBook).Methods("DELETE", "OPTIONS")

	return router
}
