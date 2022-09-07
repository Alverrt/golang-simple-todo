package main

import (
	"fmt"

	"github.com/Alverrt/golang-simple-todo/db"
	"github.com/Alverrt/golang-simple-todo/handlers"
	"github.com/Alverrt/golang-simple-todo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Connect()
	db.AutoMigrate(&models.Todo{})

	r := gin.Default()
	r.GET("/ping", handlers.StatusCheckHandler)
	r.GET("/todo", handlers.GetAllTodoItems)
	r.GET("/todo/:id", handlers.GetTodoItem)

	r.POST("/todo", handlers.InsertNewTodoItem)

	r.DELETE("/todo/:id", handlers.DeleteTodoItem)
	r.Run()

	fmt.Println("Server is running!")
}
