package usecase

import (
	"code-competence/models"
	"code-competence/models/payload"
	"code-competence/repository/database"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetListItems() (items []models.Item, err error) {
	items, err = database.GetItems()
	if err != nil {
		fmt.Println("GetListItems: Error getting list items from database")
		return
	}
	return
}

func GetItem(ID uuid.UUID) (item models.Item, err error) {
	item, err = database.GetItemById(ID)
	if err != nil {
		fmt.Println("GetUser: Error getting item by id from database")
		return
	}
	return
}

func CreateItem(c echo.Context, req *payload.CreateItemRequest) (item models.Item, err error) {

	item = models.Item{
		ID:          uuid.New(),
		Name:        req.Name,
		Category_id: req.Category_id,
		Description: req.Description,
		Stock:       req.Stock,
	}

	err = database.CreateItem(item)
	if err != nil {
		return
	}

	itemResult, err := database.GetItemById(item.ID)
	if err != nil {
		return
	}

	return itemResult, err
}

func UpdateItem(item *models.Item) (err error) {
	err = database.UpdateItem(item)
	if err != nil {
		fmt.Println("UpdateItem : Error updating item, err: ", err)
		return
	}
	return
}

func DeleteItem(id uuid.UUID) (err error) {
	item := models.Item{}
	item.ID = id
	err = database.DeleteItem(&item)
	if err != nil {
		fmt.Println("DeleteItem : error deleting item, err: ", err)
		return
	}

	return
}

func GetItemByCategory(category_id uint64) (item []models.Item, err error) {
	item, err = database.GetItemByCategoryId(category_id)
	if err != nil {
		fmt.Println("GetUser: Error getting item by category id from database")
		return
	}
	return
}

func SearchItemByName(name string) (item []models.Item, err error) {
	item, err = database.SearchItemByName(name)
	if err != nil {
		fmt.Println("GetUser: Error getting item by name from database")
		return
	}
	return
}