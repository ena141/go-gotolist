package routes

import (
	"github.com/ena141/go-todolist/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("statics/*")

	router.GET("/", controllers.GetTasks)
	router.POST("/add", controllers.AddTask)
	router.GET("/complete/:id", controllers.CompleteTask)
	router.GET("/delete/:id", controllers.DeleteTask)

	return router
}
