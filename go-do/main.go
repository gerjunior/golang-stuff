package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetAllTasksHandler(ctx *gin.Context) {
	var TaskModel Task

	tasks := TaskModel.GetAllTasks()

	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	TaskModel := &Task{Model: gorm.Model{ID: uint(id)}}
	task := TaskModel.GetOneTask()

	ctx.JSON(http.StatusOK, task)
}

func CreateTaskHandler(ctx *gin.Context) {
	var task Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.CreateTask()

	ctx.JSON(http.StatusCreated, task)
}

func UpdateTaskHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	parsedId := uint(id)

	updatedTask := Task{Model: gorm.Model{ID: parsedId}}

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedTask.UpdateTask()

	ctx.JSON(http.StatusOK, updatedTask)
}

func DeleteTaskHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}

	task := Task{Model: gorm.Model{
		ID: uint(id),
	}}

	task.DeleteTask()

	ctx.Status(http.StatusNoContent)
}

func main() {
	router := gin.Default()

	router.GET("/task", GetAllTasksHandler)
	router.GET("/task/:id", GetTaskHandler)
	router.POST("/task", CreateTaskHandler)
	router.PUT("/task/:id", UpdateTaskHandler)
	router.DELETE("/task/:id", DeleteTaskHandler)

	router.Run(":8080")
}
