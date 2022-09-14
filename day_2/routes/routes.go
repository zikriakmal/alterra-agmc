package routes

import (
	"day_2/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	//task 1 build hello world controller from root route
	e.GET("/", controllers.HelloController)
	//task 2 create status controller from status prefix
	e.GET("/status", controllers.StatusController)
	//task 3 Grouping on echo framework and implement get by param, query and post with raw json
	v1 := e.Group("/v1")

	/** for static route **/
	v1Books := v1.Group("/books")
	//get all
	v1Books.GET("", controllers.GetBooksController)
	//create
	v1Books.POST("", controllers.CreateBookController)
	//show
	v1Books.GET("/:id", controllers.GetBookByIdController)
	//destroy
	v1Books.DELETE("/:id", controllers.DeleteBookController)
	//update
	v1Books.PUT("/:id", controllers.UpdateBookController)

	/** for dynamic route **/
	v1Users := v1.Group("/users")
	//get all
	v1Users.GET("", controllers.GetUsersController)
	//create
	v1Users.POST("", controllers.CreateUserController)
	//show
	v1Users.GET("/:id", controllers.GetUserByIdController)
	//destroy
	v1Users.DELETE("/:id", controllers.DeleteUserController)
	//update
	v1Users.PUT("/:id", controllers.UpdateUserController)

	return e
}
