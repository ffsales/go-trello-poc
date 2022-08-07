package lists

import "github.com/go-chi/chi"

func GetRoutes(router *chi.Mux) {
	router.Post("/go-trello/lists", CreateList)
	router.Get("/go-trello/lists", GetAllLists)
	router.Get("/go-trello/boards/{boardId}/lists", GetListsByBoard)
	router.Get("/go-trello/lists/{listId}", GetList)
	router.Put("/go-trello/lists/{listId}", UpdateList)
	router.Delete("/go-trello/lists/{listId}", DeleteList)
}
