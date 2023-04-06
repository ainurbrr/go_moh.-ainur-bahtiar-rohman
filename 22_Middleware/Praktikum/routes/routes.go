package routes

import (
	"praktikum/constants"
	"praktikum/controllers"
	"praktikum/middlewares"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// Route / to handler function
	middlewares.LogMiddleware(e)
	e.POST("/users/add", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)

	eUsersJwt := e.Group("/users")
	eUsersJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUsersJwt.GET("", controllers.GetUsersController)
	eUsersJwt.GET("/:id", controllers.GetUserController)
	eUsersJwt.DELETE("/:id", controllers.DeleteUserController)
	eUsersJwt.PUT("/:id", controllers.UpdateUserController)

	eBooksJwt := e.Group("/books")
	eBooksJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBooksJwt.GET("", controllers.GetBooksController)
	eBooksJwt.GET("/:id", controllers.GetBookController)
	eBooksJwt.POST("", controllers.CreateBookController)
	eBooksJwt.DELETE("/:id", controllers.DeleteBookController)
	eBooksJwt.PUT("/:id", controllers.UpdateBookController)

	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)

	return e
}
