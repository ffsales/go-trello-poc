package models

import "net/http"

type Card struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
	IdList   int64  `json:"id_list"`
}

type CardResponse struct {
	*Card
}

func (*CardResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (card *Card) ToResponse() *CardResponse {
	return &CardResponse{Card: card}
}
