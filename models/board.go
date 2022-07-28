package models

type Board struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Lists       []*List `json:"lists,omitempty"`
}
