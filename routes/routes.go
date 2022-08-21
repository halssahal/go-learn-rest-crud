package routes

import (
	"learn/crud/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": http.StatusNotFound, "message": "Route not found"})
	})

	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	return r
}
