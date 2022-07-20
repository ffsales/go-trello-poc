package models

type Board struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Lists       List   `json:"Lists"`
}
