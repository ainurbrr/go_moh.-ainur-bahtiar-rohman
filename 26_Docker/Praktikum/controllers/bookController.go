package controllers

import (
	"net/http"

	"praktikum/config"
	"praktikum/models"
	"praktikum/utils"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	BookId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}
	// your solution here
	var books []models.Book

	if err := config.DB.First(&books, BookId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get books by id",
		"books":   books,
	})

}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	BookId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}
	// your solution here
	var books []models.Book

	if err := config.DB.Delete(&books, BookId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	BookId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}

	book := models.Book{}
	if err := config.DB.First(&book, BookId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Bind(&book)

	if err := config.DB.Model(&book).Updates(book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id",
		"book":    book,
	})
}
