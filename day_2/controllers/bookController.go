package controllers

import (
	"day_2/lib/database"
	"day_2/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": books,
		},
	)
}

func CreateBookController(c echo.Context) error {
	var book models.Book
	//err := c.Bind(&book)
	book = models.Book{
		ID:     1,
		Name:   "Book 1",
		Author: "Author 1",
	}

	bookCreated, err := database.CreateBook(book)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]any{
		"code": 200,
		"data": bookCreated,
	})
}

func GetBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := database.GetBookById(id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": book,
		})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := database.DeleteBookById(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": book,
		},
	)
}

func UpdateBookController(c echo.Context) error {
	var book models.Book
	id, _ := strconv.Atoi(c.Param("id"))
	//err := c.Bind(&book)

	book = models.Book{
		ID:     id,
		Name:   "Book " + strconv.Itoa(id),
		Author: "Author " + strconv.Itoa(id),
	}

	bookData, err := database.UpdateBookById(id, book)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]any{
		"code": 200,
		"data": bookData,
	})
}
