package main

import (
	"fmt"
	"log"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hi there %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
