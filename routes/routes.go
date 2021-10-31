package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/controllers/auth"
	"github.com/mellowdevs/mellow-done/controllers/homepage"
	"github.com/mellowdevs/mellow-done/controllers/list"
	"github.com/mellowdevs/mellow-done/controllers/task"
	"github.com/mellowdevs/mellow-done/middleware"
)

func InitRouter(router *gin.Engine) {

	homepageRoutes := router.Group("")
	{
		homepageRoutes.GET("", func(c *gin.Context) {
			homepage.GetHomepage(c)
		})
	}
	authRoutes := router.Group("")
	{
		authRoutes.POST("api/v1/auth/login", func(c *gin.Context) {
			auth.Login(c)
		})
		authRoutes.POST("api/v1/auth/register", func(c *gin.Context) {
			auth.Register(c)
		})
		authRoutes.GET("api/v1/auth", func(c *gin.Context) {
			auth.GetAuthenticatedUser(c)
		})
		authRoutes.GET("api/v1/auth/logout", func(c *gin.Context) {
			auth.Logout(c)
		})

	}

	listRoutes := router.Group("")
	{
		listRoutes.POST("api/v1/list", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.CreateList(c)
		})
		listRoutes.PUT("api/v1/list", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.UpdateList(c)
		})
		listRoutes.DELETE("api/v1/list", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.DeleteList(c)
		})
		listRoutes.GET("api/v1/list/id", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.GetList(c)
		})
		listRoutes.GET("api/v1/list", middleware.AuthMiddleware(), func(c *gin.Context) {
			list.GetAllLists(c)
		})
	}

	taskRoutes := router.Group("")
	{
		taskRoutes.POST("api/v1/task", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.CreateTask(c)
		})
		taskRoutes.PUT("api/v1/task", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.UpdateTask(c)
		})
		taskRoutes.DELETE("api/v1/task", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.DeleteTask(c)
		})
		taskRoutes.GET("api/v1/task/id", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.GetTaskById(c)
		})
		taskRoutes.GET("/listId", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.GetTasksByListId(c)
		})
		taskRoutes.GET("api/v1/task", middleware.AuthMiddleware(), func(c *gin.Context) {
			task.GetAllTasks(c)
		})

	}
}
