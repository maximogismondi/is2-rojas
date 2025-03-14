package model

type Item struct {
	Name  string  `json:"name" binding:"required"`
	ID    string  `json:"id" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type Error struct {
	Message string `json:"message"`
}
