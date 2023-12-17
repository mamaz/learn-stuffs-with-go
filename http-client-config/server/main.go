package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	f, err := os.OpenFile("log-server.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {

		raw, err := io.ReadAll(r.Body)
		if err != nil {
			render.JSON(w, r, map[string]any{
				"error": fmt.Sprintf("%v", err),
			})
		}

		var data map[string]any
		err = json.Unmarshal(raw, &data)
		if err != nil {
			render.JSON(w, r, map[string]any{
				"error": fmt.Sprintf("%v", err),
			})
		}

		defer r.Body.Close()

		render.JSON(w, r, map[string]any{
			"status": "ok",
			"data":   data,
		})
	})

	port := "8000"

	log.Printf("listening to port %v", port)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
		Handler:      r,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("error on estabilshing server at port: %v, error: %+v\n", port, err)
	}
}
