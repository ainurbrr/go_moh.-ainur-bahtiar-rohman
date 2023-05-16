package usecase

import (
	"code-competence/models"
	"code-competence/repository/database"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func GetListUsers() (users []models.User, err error) {
	users, err = database.GetUsers()
	if err != nil {
		fmt.Println("GetListUsers: Error getting users from database")
		return
	}
	return
}


func GetUser(ID uuid.UUID) (user *models.User, err error) {
	user, err = database.GetUserById(ID)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}

	return
}

func UpdateUser(user *models.User) (err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(passwordHash)
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}

	return
}

func DeleteUser(id uuid.UUID) (err error) {
	user := models.User{}
	user.ID = id
	err = database.DeleteUser(&user)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err: ", err)
		return
	}

	return
}
