package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]any{
			"status": "ok",
			"data":   "world",
		})
	})

	port := "3000"

	log.Printf("listening to port %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
