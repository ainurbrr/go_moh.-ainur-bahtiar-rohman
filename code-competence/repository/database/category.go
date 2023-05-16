package database

import (
	"code-competence/config"
	"code-competence/models"
)

func CreateCategory(category *models.Category) error {
	if err := config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func GetCategorys() (categorys []models.Category, err error) {
	if err = config.DB.Find(&categorys).Error; err != nil {
		return
	}
	return
}

func UpdateCategory(category *models.Category) error {
	if err := config.DB.Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategory(category *models.Category) error {
	if err := config.DB.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
