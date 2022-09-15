package controllers

import (
	"day_2/lib/database"
	"day_2/middlewares"
	"day_2/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	id := middlewares.ExtractTokenUserId(c)

	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"id":   id,
			"code": 200,
			"data": users,
		},
	)
}

func CreateUserController(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)

	userDb, err := database.CreateUser(user)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]any{
		"code": 200,
		"data": userDb,
	})
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := database.GetUserById(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": user,
		})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := database.DeleteUserById(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": user,
		},
	)
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	err := c.Bind(&user)

	userDb, err := database.UpdateUser(id, user)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]any{
		"code": 200,
		"data": userDb,
	})
}

func LoginUserController(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	user, err = database.LoginUser(user)
	token, err := middlewares.CreateToken(user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Token = token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": user,
	})

}
