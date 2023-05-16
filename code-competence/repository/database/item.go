package database

import (
	"code-competence/config"
	"code-competence/models"

	"github.com/google/uuid"
)

func GetItems() (items []models.Item, err error) {
	if err = config.DB.Model(&models.Item{}).Find(&items).Error; err != nil {
		return
	}
	return
}

func GetItemById(id uuid.UUID) (models.Item, error){
	var item models.Item

	if err := config.DB.Preload("Category").Where("id = ?", id).Find(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func CreateItem(item models.Item) (err error) {
	if err = config.DB.Create(&item).Error; err != nil {
		return
	}
	return nil
}

func UpdateItem(item *models.Item) error {
	if err := config.DB.Model(&item).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(item *models.Item) error {
	if err := config.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}

func GetItemByCategoryId(category_id uint64) ([]models.Item, error){
	var item []models.Item

	if err := config.DB.Preload("Category").Where("category_id = ?", category_id).Find(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func SearchItemByName(name string) ([]models.Item, error){
	var item []models.Item

	name = "%"+name+"%"

	if err := config.DB.Preload("Category").Where("name LIKE ?", name).Find(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

