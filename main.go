package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

var db mongoDB

func main() {
	fmt.Println("Golang+GraphQL+MongoDB Server v1.0.0")

	// MongoDB
	db = connectDB()
	defer db.closeDB()

	// GraphQL
	schema := initSchema()
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // GraphiQL interface
	})

	// Serve
	http.Handle("/", http.FileServer(http.Dir("./public"))) // Serve the frontend in /public
	http.Handle("/graphql", disableCors(headerAuthorization(h)))
	err := http.ListenAndServe(config.serveUri, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
