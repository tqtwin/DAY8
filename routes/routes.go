// routes/routes.go
package routes

import (
	"day8/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, redisCli *redis.Client) {
	userController := controllers.NewUserController(db, redisCli)
	postController := controllers.NewPostController(db, redisCli) // Khởi tạo PostController

	apiGroup := router.Group("/api")
	{
		// User routes
		apiGroup.GET("/users", userController.GetUsers)
		apiGroup.POST("/users", userController.CreateUser)
		apiGroup.GET("/users/:id", userController.GetUserByID)
		apiGroup.PUT("/users/:id", userController.UpdateUser)
		apiGroup.DELETE("/users/:id", userController.DeleteUser)

		// Post routes
		apiGroup.GET("/posts", postController.GetPosts)
		apiGroup.POST("/posts", postController.CreatePost)
		apiGroup.GET("/posts/:id", postController.GetPostByID)
		apiGroup.PUT("/posts/:id", postController.UpdatePost)
		apiGroup.DELETE("/posts/:id", postController.DeletePost)
	}
}
