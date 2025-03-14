package controller

import (
	"net/http"
	"rest-workshop/ej2/src/model"
	"rest-workshop/ej2/src/service"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	service service.ItemService
}

func (i *ItemController) AddItem(c *gin.Context) {
	var item model.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: err.Error()})
		return
	} else if _, err := i.service.Create(item); err != nil {
		c.JSON(http.StatusConflict, model.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)

}

func NewItemController(service service.ItemService) ItemController {
	return ItemController{service}
}
