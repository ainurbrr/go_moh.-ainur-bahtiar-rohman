package controllers

import (
	"code-competence/models"
	"code-competence/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategorysController(c echo.Context) error {
	categorys, e := usecase.GetListCategorys()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"Categorys": categorys,
	})
}

// create new Category
func CreateCategoryController(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)

	if err := usecase.CreateCategory(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create Category",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new Category",
		"Category": category,
	})
}

// delete Category by id
func DeleteCategoryController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteCategory(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete Category",
			"errorDescription": err,
			"errorMessage":     "Sorry cannot delete category",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Category",
	})
}

// update Category by id
func UpdateCategoryController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	category := models.Category{}
	c.Bind(&category)
	category.ID = uint(id)
	if err := usecase.UpdateCategory(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update Category",
			"errorDescription": err,
			"errorMessage":     "Sorry cannot update category",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Category",
	})
}
