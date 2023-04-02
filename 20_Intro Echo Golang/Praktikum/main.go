package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userId, err := getIdFromParam(c)
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Id == userId {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get selected users by id",
				"users":    user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userId, err := getIdFromParam(c)
	if err != nil {
		return err
	}

	for i, user := range users {
		if user.Id == userId {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete users by id: " + strconv.Itoa(userId),
				"users":    users,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	userId, err := getIdFromParam(c)
	if err != nil {
		return err
	}
	updatedUser := User{}
	c.Bind(&updatedUser)
	updatedUser.Id = userId

	for i, user := range users {
		if user.Id == userId {
			users[i] = updatedUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update users by id: " + strconv.Itoa(userId),
				"users":    users,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func getIdFromParam(e echo.Context) (int, error) {
	id, err := strconv.Atoi(e.Param("id"))
	return id, err
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
