package models

type List struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	IdBoard int    `json:"id_board"`
}
