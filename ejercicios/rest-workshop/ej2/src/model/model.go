package model

type Item struct {
	Name  string `json:"name" binding:"required"`
	ID    string `json:"id" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
