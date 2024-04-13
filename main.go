package main

import (
	"net/http"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.TestRoute(r)
	http.ListenAndServe("127.0.0.1:3000", r)
}