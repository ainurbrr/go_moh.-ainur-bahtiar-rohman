package usecase

import (
	"code-competence/middlewares"
	"code-competence/models"
	"code-competence/models/payload"
	"code-competence/repository/database"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	IDuuid := uuid.New()

	if !database.IsEmailAvailable(req.Email) {
		return resp, errors.New("email is not available")
	}

	newUser := &models.User{
		ID:       IDuuid,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passwordHash),
	}

	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	token, err := middlewares.CreateToken(IDuuid)
	if err != nil {
		fmt.Println("Error creating token")
	}

	resp = payload.CreateUserResponse{
		UserID: newUser.ID,
		Name:   newUser.Name,
		Token:  token,
	}

	return
}

func LoginUser(c echo.Context, req *payload.LoginRequest) (resp payload.LoginResponse, err error) {
	user, err := database.LoginUser(req.Email)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database/User not found")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return
	}
	token, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return resp, err
	}
	resp = payload.LoginResponse{
		UserID: user.ID,
		Name:   user.Name,
		Token:  token,
	}
	return resp, nil

}
