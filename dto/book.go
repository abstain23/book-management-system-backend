package dto

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}
