package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	Description string
}

func init() {
	db = GetDB()
	db.AutoMigrate(&Task{})
}

func (t *Task) GetAllTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func (t Task) GetOneTask() Task {
	var tasks []Task
	db.First(&tasks, t)
	fmt.Println(tasks[0])
	return tasks[0]
}

func (t *Task) CreateTask() *Task {
	db.NewRecord(t)
	db.Create(t)
	return t
}

func (t *Task) UpdateTask() *Task {
	db.Save(&t)
	return t
}

func (t *Task) DeleteTask() *Task {
	db.Delete(&t)
	return t
}
