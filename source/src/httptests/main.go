package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Item Created"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", CreateEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe("https://www.alexedwards.net/blog/golang-response-snippets", router))
}
