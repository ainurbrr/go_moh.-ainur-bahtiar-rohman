package usecase

import (
	"code-competence/models"
	"code-competence/repository/database"
	"fmt"
)

func CreateCategory(category *models.Category) error {

	err := database.CreateCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func GetListCategorys() (categorys []models.Category, err error) {
	categorys, err = database.GetCategorys()
	if err != nil {
		fmt.Println("GetListCategorys: Error getting Categorys from database")
		return
	}
	return
}

func UpdateCategory(category *models.Category) (err error) {
	err = database.UpdateCategory(category)
	if err != nil {
		fmt.Println("UpdateCategory : Error updating category, err: ", err)
		return
	}

	return
}

func DeleteCategory(id uint) (err error) {
	category := models.Category{}
	category.ID = id
	err = database.DeleteCategory(&category)
	if err != nil {
		fmt.Println("DeleteCategory : error deleting Category, err: ", err)
		return
	}

	return
}
