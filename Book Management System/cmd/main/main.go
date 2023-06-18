package main

import (
	"book-management-system/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}