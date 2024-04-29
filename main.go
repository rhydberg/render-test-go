package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
func api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API"))
}

func main(){
	http.HandleFunc("/", hello)
	http.HandleFunc("/api", api)
	log.Fatal(http.ListenAndServe(":8080", nil))
}