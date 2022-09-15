package database

import (
	"day_2/models"
	"strconv"
)

func GetBooks() ([]models.Book, error) {
	books := []models.Book{
		{
			ID:     1,
			Name:   "Book 1",
			Author: "Author 1",
		},
		{
			ID:     2,
			Name:   "Book 2",
			Author: "Author 2",
		},
	}
	return books, nil
}

func GetBookById(id int) (models.Book, error) {
	book := models.Book{
		ID:     id,
		Name:   "Book " + strconv.Itoa(id),
		Author: "Author " + strconv.Itoa(id),
	}
	return book, nil
}

func CreateBook(book models.Book) (models.Book, error) {
	book = models.Book{
		ID:     1,
		Name:   book.Name,
		Author: book.Author,
	}
	return book, nil
}

func UpdateBookById(id int, book models.Book) (models.Book, error) {
	bookUpdate := models.Book{
		ID:     id,
		Name:   book.Name,
		Author: book.Author,
	}
	return bookUpdate, nil
}

func DeleteBookById(id int) (interface{}, error) {
	return map[string]string{}, nil
}
