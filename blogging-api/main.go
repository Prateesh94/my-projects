package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/blogops"
	"main.go/clientandcache"
)

func main() {
	router := mux.NewRouter()
	router.Use(clientandcache.Limitmid)
	router.HandleFunc("/add", blogops.AddData).Methods("PUT")
	router.HandleFunc("/delete/{id}", blogops.DelData).Methods("DELETE")
	router.HandleFunc("/view/{id}", blogops.View).Methods("GET")
	router.HandleFunc("/view/", blogops.View).Methods("GET")
	router.HandleFunc("/view", blogops.View).Methods("GET")
	router.HandleFunc("/update/{id}", blogops.UpData).Methods("POST")
	router.HandleFunc("/search/{term}", blogops.Blogsearch).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
	
}
