package models

import "net/http"

type Board struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Lists       []*List `json:"lists,omitempty"`
}

type BoardResponse struct {
	*Board
}

func (*BoardResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (board *Board) ToResponse() *BoardResponse {
	return &BoardResponse{Board: board}
}
