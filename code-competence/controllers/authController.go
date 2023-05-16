package controllers

import (
	"code-competence/models/payload"
	"code-competence/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	var payloadLogin payload.LoginRequest
	c.Bind(&payloadLogin)

	if err := c.Validate(payloadLogin); err != nil {
		return err
	}

	user, err := usecase.LoginUser(c, &payloadLogin)
	c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})

}
