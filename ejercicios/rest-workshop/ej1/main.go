package main

import (
	"github.com/gin-gonic/gin"
)

type Item struct {
	Name  string `json:"name" binding:"required"`
	ID    string `json:"id" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type Error struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	items := make(map[string]Item, 0)

	r.POST("/items", func(c *gin.Context) {
		var payload Item

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, Error{Message: err.Error()})
			return
		}

		if _, found := items[payload.ID]; found {
			c.JSON(409, Error{Message: "Item already exists"})
			return
		}

		items[payload.ID] = payload
		c.JSON(201, payload)
	})

	r.Run()
}
