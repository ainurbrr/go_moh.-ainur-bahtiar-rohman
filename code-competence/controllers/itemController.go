package controllers

import (
	"code-competence/models"
	"code-competence/models/payload"
	"code-competence/usecase"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetItemsController(c echo.Context) error {
	items, err := usecase.GetListItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "succes",
		"items":  items,
	})

}

func GetItemController(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))
	item, err := usecase.GetItem(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get item",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"item":   item,
	})
}

func CreateItemController(c echo.Context) error {
	payloadItem := payload.CreateItemRequest{}
	c.Bind(&payloadItem)

	if err := c.Validate(payloadItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create item",
			"error":    err.Error(),
		})
	}

	item, err := usecase.CreateItem(c, &payloadItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error create item",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"data":    item,
	})
}

func UpdateItemController(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))
	item := models.Item{}
	c.Bind(&item)
	item.ID = id
	if err := usecase.UpdateItem(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update item",
			"errorDescription": err,
			"errorMessage":     "Sorry cannot update item",
		})
	}

	itemUpdated, _ := usecase.GetItem(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success update item",
		"updated item": itemUpdated,
	})
}

func DeleteItemController(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	if err := usecase.DeleteItem(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete item",
			"errorDescription": err,
			"errorMessage":     "Sorry cannot delete item",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item",
	})
}

func GetCategoryItemController(c echo.Context) error {
	category_id, _ := strconv.ParseUint(c.Param("category_id"),10, 32)
	CategoryItem, err := usecase.GetItemByCategory(category_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get item",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"item":   CategoryItem,
	})
}


func SearcItemNameController(c echo.Context) error {
	name := c.QueryParam("keyword")
	itemByName, err := usecase.SearchItemByName(name)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get item",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"item":   itemByName,
	})
}
