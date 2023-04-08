package routes

import (
	"praktikum/constants"
	"praktikum/controllers"
	"praktikum/middlewares"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// Route / to handler function
	middlewares.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	eUsersJwt := e.Group("/users")
	eUsersJwt.GET("", controllers.GetUsersController, mid.JWT([]byte(constants.SECRET_JWT)))
	eUsersJwt.GET("/:id", controllers.GetUserController, mid.JWT([]byte(constants.SECRET_JWT)))
	eUsersJwt.POST("", controllers.CreateUserController)
	eUsersJwt.POST("/login", controllers.LoginUserController)
	eUsersJwt.DELETE("/:id", controllers.DeleteUserController, mid.JWT([]byte(constants.SECRET_JWT)))
	eUsersJwt.PUT("/:id", controllers.UpdateUserController, mid.JWT([]byte(constants.SECRET_JWT)))

	eBooksJwt := e.Group("/books")
	eBooksJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBooksJwt.GET("", controllers.GetBooksController)
	eBooksJwt.GET("/:id", controllers.GetBookController)
	eBooksJwt.POST("", controllers.CreateBookController)
	eBooksJwt.DELETE("/:id", controllers.DeleteBookController)
	eBooksJwt.PUT("/:id", controllers.UpdateBookController)

	eBlogsJwt := e.Group("/blogs")
	eBlogsJwt.GET("", controllers.GetBlogsController)
	eBlogsJwt.GET("/:id", controllers.GetBlogController)
	eBlogsJwt.POST("", controllers.CreateBlogController)
	eBlogsJwt.DELETE("/:id", controllers.DeleteBlogController)
	eBlogsJwt.PUT("/:id", controllers.UpdateBlogController)

	return e
}
