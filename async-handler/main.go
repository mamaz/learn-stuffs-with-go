package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-go/handler"

	"github.com/go-chi/chi"

	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", handler.Handle())

	port := "3000"

	log.Printf("listening to port %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
