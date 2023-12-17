package main

import (
	"bytes"
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
	f, err := os.OpenFile("log-client.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	sender := Sender{}

	// transport := http.DefaultTransport.(*http.Transport).Clone()
	// transport.MaxIdleConns = 1000
	// transport.MaxConnsPerHost = 1000
	// transport.MaxIdleConnsPerHost = 1000
	// transport.IdleConnTimeout = 100 * time.Second

	client := &http.Client{
		Timeout: 5 * time.Second,
		// Transport: transport,
	}

	/**
	&http.Transport{
			MaxIdleConns:        10,                     // In mw-gql-order: 10
			MaxIdleConnsPerHost: 10,                     // In mw-gql-order: 10
			IdleConnTimeout:     100 * time.Millisecond, // In mw-gql-order: 100 * time.Milisecond
			MaxConnsPerHost:     10,                     // In mw-gql-order: NO LIMIT
		},
	*/

	log.Printf("DEFAULT TRANSPORT: %+v", client.Transport)

	URL := "http://localhost:8000/hello"

	r.Get("/send", func(w http.ResponseWriter, r *http.Request) {
		resp, err := sender.SendGET(client, "GET", URL, bytes.NewBufferString(`{"message":"hello"}`))
		if err != nil {
			log.Printf("eror on receiving value: %v", err)
		}
		defer resp.Body.Close()

		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			render.JSON(w, r, map[string]any{
				"error": fmt.Sprintf("%v", err),
			})
		}

		var data map[string]any
		json.Unmarshal(raw, &data)

		render.JSON(w, r, data)
	})

	port := "9000"

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

type Sender struct {
}

func (s *Sender) SendGET(client *http.Client, method string, URL string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
