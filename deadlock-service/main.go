package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"deadlock-service/service"
)

type InsertRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (ir *InsertRequest) Bind(r *http.Request) error {
	return nil
}

type ErrRender struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Render implements render.Renderer
func (er ErrRender) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusInternalServerError)
	return nil
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"ok"`
}

func (resp Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusAccepted)
	return nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	r.Post("/data", func(w http.ResponseWriter, r *http.Request) {
		var request InsertRequest
		err := render.Bind(r, &request)
		if err != nil {
			log.Printf("error on binding data: %+v\n", err)
			render.Render(w, r, ErrRender{Status: "error", Message: "internal server error"})
			return
		}

		result := service.UpdateData(request.ID, request.Name)

		log.Printf("data is received and result is: %+v\n", result)
		render.Render(w, r, Response{Status: "accepted", Message: "request has been accepted"})
	})

	http.ListenAndServe(":3000", r)
}
