package routes

import (
	"code-competence/constants"
	"code-competence/controllers"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Routes(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(mid.CORS())
	e.Pre(mid.RemoveTrailingSlash())

	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.RegisterUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserController)
	user.POST("/register", controllers.RegisterUserController)
	user.PUT("/:id", controllers.UpdateUserController, mid.JWT([]byte(constants.SECRET_JWT)))
	user.DELETE("/:id", controllers.DeleteUserController, mid.JWT([]byte(constants.SECRET_JWT)))

	category := e.Group("/category")
	category.GET("", controllers.GetCategorysController)
	category.POST("", controllers.CreateCategoryController)
	category.PUT("/:id", controllers.UpdateCategoryController)
	category.DELETE("/:id", controllers.DeleteCategoryController)

	item := e.Group("/items")
	item.GET("", controllers.GetItemsController, mid.JWT([]byte(constants.SECRET_JWT)))
	item.GET("/:id", controllers.GetItemController, mid.JWT([]byte(constants.SECRET_JWT)))
	item.POST("", controllers.CreateItemController, mid.JWT([]byte(constants.SECRET_JWT)))
	item.PUT("/:id", controllers.UpdateItemController, mid.JWT([]byte(constants.SECRET_JWT)))
	item.DELETE("/:id", controllers.DeleteItemController, mid.JWT([]byte(constants.SECRET_JWT)))
	item.GET("/category/:category_id", controllers.GetCategoryItemController, mid.JWT([]byte(constants.SECRET_JWT)))
	item.GET("/category", controllers.SearcItemNameController, mid.JWT([]byte(constants.SECRET_JWT)))

}
