package main

import (
	"net/http"
	"io"
	"log"
) 

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"hello":"world"}`)
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("error on listening", err)
	}
}
