package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/prakash-joshi/go-postgres-gorm/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	fmt.Println("Starting server on Port 8080....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
