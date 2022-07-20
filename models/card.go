package models

type Card struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
}
