package cards

import "github.com/go-chi/chi"

func GetRoutes(router *chi.Mux) {
	router.Post("/go-trello/cards", CreateCard)
	router.Get("/go-trello/cards", GetAllCards)
	router.Get("/go-trello/lists/{listId}/cards", GetCardsByList)
	router.Get("/go-trello/cards/{cardId}", GetCard)
	router.Put("/go-trello/cards/{cardId}", UpdateCard)
	router.Delete("/go-trello/cards/{cardId}", DeleteCard)
}
