package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RightTime(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", timeNow)

}

func main() {
	http.HandleFunc("/", RightTime)
	log.Println("Executando....")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
