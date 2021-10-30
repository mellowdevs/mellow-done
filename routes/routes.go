package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/controllers/auth"
	"github.com/mellowdevs/mellow-done/controllers/list"
	"github.com/mellowdevs/mellow-done/controllers/task"
	"github.com/mellowdevs/mellow-done/middleware"
)

func InitRouter(router *gin.Engine) {
	authRoutes := router.Group("api/v1/auth")
	{
		authRoutes.POST("/login", func(c *gin.Context) {
			auth.Login(c)
		})
		authRoutes.POST("/register", func(c *gin.Context) {
			auth.Register(c)
		})
		authRoutes.GET("", func(c *gin.Context) {
			auth.GetAuthenticatedUser(c)
		})
		authRoutes.GET("/logout", func(c *gin.Context) {
			auth.Logout(c)
		})

	}

	listRoutes := router.Group("api/v1/list")
	{
		listRoutes.POST("", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.CreateList(c)
		})
		listRoutes.PUT("", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.UpdateList(c)
		})
		listRoutes.DELETE("", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.DeleteList(c)
		})
		listRoutes.GET("/id", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.GetList(c)
		})
		listRoutes.GET("", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.GetAllLists(c)
		})
	}

	taskRoutes := router.Group("api/v1/task")
	{
		taskRoutes.POST("", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.CreateTask(c)
		})
		taskRoutes.PUT("", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.UpdateTask(c)
		})
		taskRoutes.DELETE("", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.DeleteTask(c)
		})
		taskRoutes.GET("/id", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.GetTaskById(c)
		})
		taskRoutes.GET("/listId", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.GetTasksByListId(c)
		})
		taskRoutes.GET("", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.GetAllTasks(c)
		})

	}
}
