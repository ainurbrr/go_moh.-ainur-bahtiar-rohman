package controllers

import (
	"code-competence/middlewares"
	"code-competence/models"
	"code-competence/models/payload"
	"code-competence/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := usecase.GetListUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))
	user, err := usecase.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get user",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

func RegisterUserController(c echo.Context) error {
	payloadUser := payload.CreateUserRequest{}
	c.Bind(&payloadUser)

	if err := c.Validate(payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create user",
			"error":    err.Error(),
		})
	}

	resp, err := usecase.CreateUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    resp,
	})
}

func DeleteUserController(c echo.Context) error {
	userId, _ := middlewares.ExtractTokenId(c)
	if userId == uuid.Nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": "userId not found",
			"errorMessage":     "Sorry Token is invalid",
		})
	}
	id, _ := uuid.Parse(c.Param("id"))
	if userId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized User to Delete")
	}

	if err := usecase.DeleteUser(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf user tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func UpdateUserController(c echo.Context) error {
	userId, _ := middlewares.ExtractTokenId(c)
	if userId == uuid.Nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": "userId not found",
			"errorMessage":     "Sorry Token is invalid",
		})
	}
	id, _ := uuid.Parse(c.Param("id"))
	if userId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized User to update")
	}
	user := models.User{}
	c.Bind(&user)
	user.ID = id
	if err := usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update user",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf user tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
