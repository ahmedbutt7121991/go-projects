package main

import (
	"log"
	"net/http"

	"github.com/ahmedbutt7121991/go-projects/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", r))
}
