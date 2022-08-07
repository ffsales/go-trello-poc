package models

import "net/http"

type List struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	IdBoard int    `json:"id_board"`
}

type ListResponse struct {
	*List
}

func (*ListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (list *List) ToResponse() *ListResponse {
	return &ListResponse{List: list}
}
