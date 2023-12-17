package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Handler struct {
}

// DoHandle handle yang langsung return "ok" sembari ada goroutine yang jalanin string FINISH
// setelah sleep
func (g *Handler) DoHandle() string {
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("FINISH!")
	}()
	return "ok"
}

func Handle() func(w http.ResponseWriter, r *http.Request) {
	handler := Handler{}

	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]any{
			"status": "ok",
			"data":   handler.DoHandle(),
		})
	}
}
