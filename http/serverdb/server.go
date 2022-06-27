package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", UserHandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
