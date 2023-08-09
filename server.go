package main

import (
	"fmt"
	h "https://github.com/nwoik/Duel-Masters-DB/handlers"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/home", h.HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
