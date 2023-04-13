package controllers

import (
	"net/http"

	"praktikum/config"
	"praktikum/models"
	"praktikum/utils"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get book by id
func GetBlogController(c echo.Context) error {
	BlogId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}
	// your solution here
	blogs := models.Blog{}
	var user []models.User

	if err := config.DB.First(&blogs, BlogId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Find(&user, blogs.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// if err := config.DB.Preload("Blog").Find(&user, blogs.UserID).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blog":    blogs,
		"writer":  user,
	})

}

// create new book
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete book by id
func DeleteBlogController(c echo.Context) error {
	// your solution here
	BlogId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}
	// your solution here
	var blogs []models.Blog

	if err := config.DB.Delete(&blogs, BlogId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog by id",
	})
}

// update book by id
func UpdateBlogController(c echo.Context) error {
	// your solution here
	BlogId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}

	blog := models.Blog{}
	if err := config.DB.First(&blog, BlogId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Bind(&blog)

	if err := config.DB.Model(&blog).Updates(blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog by id",
		"blog":    blog,
	})
}
