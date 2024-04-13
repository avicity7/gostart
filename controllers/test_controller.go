package controllers

import (
	"net/http"
	"server/services"

	"github.com/go-chi/render"
)

func GetTestController(w http.ResponseWriter, r *http.Request) {
	result := services.GetTestService()
	w.WriteHeader(200)
	render.JSON(w, r, result)
}
