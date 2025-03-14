package service

import (
	"rest-workshop/ej2/src/database"
	"rest-workshop/ej2/src/model"
)

type ItemService struct {
	db database.Database[model.Item]
}

func (i ItemService) Create(item model.Item) (model.Item, error) {
	return i.db.Create(item.ID, item)
}

func NewItemService(db database.Database[model.Item]) ItemService {
	return ItemService{db}
}
