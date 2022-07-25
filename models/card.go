package models

type Card struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
	IdList   int64  `json:"id_list"`
}
