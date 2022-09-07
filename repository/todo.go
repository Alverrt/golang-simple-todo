package repository

import (
	"errors"
	"fmt"

	"github.com/Alverrt/golang-simple-todo/db"
	"github.com/Alverrt/golang-simple-todo/models"
)

func FindAll() []models.Todo {
	db := db.Connect()
	var todos []models.Todo
	result := db.Find(&todos)

	if result.Error != nil {
		fmt.Printf("There is an error: %s", result.Error.Error())
	}

	return todos
}

func FindOne(id uint) (models.Todo, int) {
	db := db.Connect()
	var todo models.Todo
	result := db.First(&todo, id)

	if result.Error != nil {
		fmt.Printf("There is an error: %s", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return todo, int(result.RowsAffected)
	}

	return todo, int(result.RowsAffected)
}

func Insert(item models.Todo) (int, error) {
	db := db.Connect()
	result := db.Create(&item)

	if result.Error != nil {
		fmt.Printf("There is an error: %s", result.Error.Error())
		return 0, errors.New(result.Error.Error())
	}

	return int(result.RowsAffected), nil
}

func Delete(id uint) (int, error) {
	db := db.Connect()
	result := db.Delete(&models.Todo{}, id)

	if result.Error != nil {
		fmt.Printf("There is an error: %s", result.Error.Error())
		return 0, errors.New(result.Error.Error())
	}

	return int(result.RowsAffected), nil
}
