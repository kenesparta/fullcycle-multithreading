package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kensparta/fullcycle-multithreading/constants"
	"github.com/kensparta/fullcycle-multithreading/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{cep}", handlers.GetCep)
	err := http.ListenAndServe(constants.Port, r)
	if err != nil {
		log.Printf("error running the server %v", err)
	}
}
