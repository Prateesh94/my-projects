package main

import (
	"log"
	"markdown-api/data"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/upload", data.UploadEndpoint)
	router.HandleFunc("/list", data.ListFilesEndpoint)
	router.HandleFunc("/view", data.CreateHTMLEndpoint)
	log.Fatal(http.ListenAndServe(":8080", router))
}
