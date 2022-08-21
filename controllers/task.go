package controllers

import (
	"learn/crud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"Task"`
	Deadline   string `json:"Deadline"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"Task"`
	Deadline   string `json:"Deadline"`
}

/*
	GET /tasks
	Get all tasks
*/
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

/*
	POST /tasks
	Create new tasks
*/
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02 15:04:05"
	deadline, _ := time.ParseInLocation(date, input.Deadline, time.Local)

	// Create task
	task := models.Task{
		AssignedTo: input.AssignedTo,
		Task:       input.Task,
		Deadline:   deadline,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

/*
	GET /tasks/:id
	Find a task
*/
func FindTask(c *gin.Context) {
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found !!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

/*
	PATCH /tasks/:id
	Update a task
*/
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get models if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found !!!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02 15:04:05"
	deadline, _ := time.ParseInLocation(date, input.Deadline, time.Local)

	var UpdateInput models.Task
	UpdateInput.Deadline = deadline
	UpdateInput.AssignedTo = input.AssignedTo
	UpdateInput.Task = input.Task

	db.Model(&task).Updates(UpdateInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

/*
	DELETE /tasks/:id
	Delete a task
*/
func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found !!!"})
		return
	}

	db.Delete(task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
