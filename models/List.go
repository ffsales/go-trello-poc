package models

type List struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Cards Card   `json:"cards"`
	Order int    `json:"order"`
}
