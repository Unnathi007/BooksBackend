package main

import (
	"fmt"
	"github.com/rs/cors"
	"goCrudDemo/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.Router()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	fmt.Println("Server at 8080")
	http.Handle("/", r)
	handler := cors.Handler(r)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("There's an error with the server", err)
	}
	fmt.Println("Server at 8080")
}
