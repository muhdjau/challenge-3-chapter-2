package controllers

import (
	"challenge-chapter-2-sesi-3/config"
	"challenge-chapter-2-sesi-3/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var bookDatas models.Books

	if err := c.ShouldBindJSON(&bookDatas); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := config.CreateBook(bookDatas)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to create book data")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Book data has created",
	})
}

func GetAllBooks(c *gin.Context) {
	var bookDatas []models.Books

	getBooks, err := config.GetAllBooks(bookDatas)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to get all books data")
		return
	}

	if getBooks == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": getBooks,
	})
}

func GetBookById(c *gin.Context) {
	bookID := c.Param("bookID")
	var bookDatas models.Books

	convBookID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Failed to convert id")
		return
	}

	book, err := config.GetBookById(convBookID, bookDatas)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %v not found", convBookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("bookID")
	var bookDatas models.Books

	convBookID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Failed to convert id")
		return
	}

	if err = c.ShouldBindJSON(&bookDatas); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = config.UpdateBook(convBookID, bookDatas)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", convBookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %d succesfully updated", convBookID),
	})
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("bookID")

	convBookID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Failed to convert id")
		return
	}

	err = config.DeleteBook(convBookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", convBookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %d succesfully deleted", convBookID),
	})
}
