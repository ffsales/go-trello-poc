package main

import (
	"fmt"
	"net/http"

	"github.com/ffsales/go-trello-poc/routes/boards"
	"github.com/ffsales/go-trello-poc/routes/cards"
	"github.com/ffsales/go-trello-poc/routes/lists"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("Iniciando o |||::Go Trello::|||")

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	boards.GetRoutes(router)
	lists.GetRoutes(router)
	cards.GetRoutes(router)

	http.ListenAndServe(":3000", router)
}
