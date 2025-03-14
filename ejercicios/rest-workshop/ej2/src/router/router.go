package router

import (
	"rest-workshop/ej2/src/controller"
	"rest-workshop/ej2/src/database"
	"rest-workshop/ej2/src/model"
	"rest-workshop/ej2/src/service"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	db := database.NewDatabase[model.Item]()
	service := service.NewItemService(db)
	controller := controller.NewItemController(service)

	r.POST("/items", controller.AddItem)

	return r
}
