package handlers

import (
	"net/http"
	"strconv"

	"github.com/Alverrt/golang-simple-todo/models"
	"github.com/Alverrt/golang-simple-todo/repository"
	"github.com/gin-gonic/gin"
)

func GetAllTodoItems(c *gin.Context) {
	result := repository.FindAll()
	c.JSON(http.StatusOK, result)
}

func GetTodoItem(c *gin.Context) {
	id := c.Params.ByName("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer."})
		return
	}

	result, count := repository.FindOne(uint(parsedId))
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, result)
}

func InsertNewTodoItem(c *gin.Context) {
	var body models.Todo

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := repository.Insert(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.JSON(http.StatusOK, true)
}

func DeleteTodoItem(c *gin.Context) {
	id := c.Params.ByName("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer."})
		return
	}

	_, err2 := repository.Delete(uint(parsedId))

	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}
