package utils

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type ResponseError struct {
	Message string `json: "message"`
}

func (*ResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Ok(w http.ResponseWriter, r *http.Request, renderer render.Renderer) {
	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
	render.Render(w, r, renderer)
}

func OkList(w http.ResponseWriter, r *http.Request, renderer []render.Renderer) {
	w.Header().Set("Content-Type", "application/json")
	render.Status(r, http.StatusOK)
	render.RenderList(w, r, renderer)
}

func NotFoundError(w http.ResponseWriter, r *http.Request, err error, message string) {
	returnError(w, r, http.StatusNotFound, err, message)
}

func UnprocessableEntityError(w http.ResponseWriter, r *http.Request, err error, message string) {
	returnError(w, r, http.StatusUnprocessableEntity, err, message)
}

func ServiceUnavailableError(w http.ResponseWriter, r *http.Request, err error, message string) {
	returnError(w, r, http.StatusServiceUnavailable, err, message)
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error, message string) {
	returnError(w, r, http.StatusBadRequest, err, message)
}

func returnError(w http.ResponseWriter, r *http.Request, status int, err error, message string) {
	log.Println("Error: ", err)
	response := &ResponseError{Message: message}
	render.Status(r, status)
	render.Render(w, r, response)
}
