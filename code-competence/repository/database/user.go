package database

import (
	"code-competence/config"
	"code-competence/models"

	"github.com/google/uuid"
)

func GetUsers() (users []models.User, err error) {
	if err = config.DB.Model(&models.User{}).Find(&users).Error; err != nil {
		return
	}
	return
}

func GetUserById(ID uuid.UUID) (user *models.User, err error) {
	user = &models.User{}
	if err := config.DB.Model(&user).Where("id = ?", ID).Find(&user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateUser(user *models.User) error {
	if err := config.DB.Model(&user).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}