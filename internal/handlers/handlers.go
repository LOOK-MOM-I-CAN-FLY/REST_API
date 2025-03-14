package handlers

import (
	"log"
	"net/http"
	"strconv"

	"compoze_rest_api/internal/db"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	log.Println("Запрос на получение всех книг")
	c.JSON(http.StatusOK, db.Db)
}
func GetBookById(c *gin.Context) {
	log.Println("Запрос на получение книги по id")
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "для этого ID не поддерживается приведение целых чисел",
		})
		return
	}
	book, ok := db.FindBookById(id)
	if ok != true {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Книга с данным ID не найдена в базе данных",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}
func CreateBook(c *gin.Context) {
	log.Println("Запрос на добавление книги")
	var book db.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Db = append(db.Db, book)
	c.JSON(http.StatusOK, book)
}
func UpdateBook(c *gin.Context) {
	log.Println("Запрос на обновление книги")
	var book db.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}
func DeleteBook(c *gin.Context) {
	log.Println("Запрос на удаление книги")
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "для этого ID не поддерживается приведение целых чисел",
		})
		return
	}
	c.JSON(http.StatusOK, db.DeleteBookById(id))
}
