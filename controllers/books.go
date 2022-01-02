package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest-go/models"
)

// Move this to a validators module
type CreateBookValidator struct {
	Title	string	`json:"title" binding:"required"`
	Author	string	`json:"author" binding:"required"`
}

type UpdateBookValidator struct {
	Title	string	`json:"title"`
	Author	string	`json:"author"`
}

func GetBooks(c *gin.Context) {
	/*
		Fetch all books and return
	*/
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}


func CreateBook(c *gin.Context) {
	/*
		Validate the input by passing the request body from the context to the input
		validator and return an error if the required fields are missing else, create
		the book and return it
	*/
	var input CreateBookValidator

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

func FetchBook(c *gin.Context) {
	/*
		Using gorm, query the DB for books that match the given ID, then get the first
		item and return it. Return an error if the book doesn't exist
	*/
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	/*
		Query the DB for a book with given ID, then validate input before updating DB.
		Return an error if either the fetch query or the validator are unsuccessful
	*/
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookValidator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fields were provided"})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	/*
		Query the DB for a book with given ID, and throw an exception if it doesn't exist.
		Delete the boook if it exists, then return a no content status code
	*/
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}