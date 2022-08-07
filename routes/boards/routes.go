package boards

import "github.com/go-chi/chi"

func GetRoutes(router *chi.Mux) {
	router.Post("/go-trello/boards", CreateBoard)
	router.Get("/go-trello/boards", GetAllBoards)
	router.Get("/go-trello/boards/{boardId}", GetBoard)
	router.Put("/go-trello/boards/{boardId}", UpdateBoard)
	router.Delete("/go-trello/boards/{boardId}", DeleteBoard)
}
