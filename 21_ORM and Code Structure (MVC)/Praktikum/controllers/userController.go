package controllers

import (
	"net/http"

	"praktikum/config"
	"praktikum/models"
	"praktikum/utils"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	UserId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}
	// your solution here
	var user []models.User
	if err := config.DB.First(&user, UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Preload("Blog").Find(&user, UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	UserId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}
	// your solution here
	var users []models.User

	if err := config.DB.Delete(&users, UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete users by id",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	UserId, err := utils.GetIdFromParam(c)
	if err != nil {
		return err
	}

	user := models.User{}
	if err := config.DB.First(&user, UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Bind(&user)

	if err := config.DB.Model(&user).Updates(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"user":    user,
	})
}
