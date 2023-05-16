package database

import (
	"code-competence/config"
	"code-competence/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(user *models.User) (err error) {
	if err = config.DB.Create(&user).Error; err != nil {
		return
	}
	return nil
}

func IsEmailAvailable(email string) bool {
	var count int64
	user := models.User{}
	if err := config.DB.Model(&user).Where("email = ?", email).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}
	return count == 0
}

func LoginUser(email string) (user models.User, err error) {
	if err = config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return
	}
	return
}
