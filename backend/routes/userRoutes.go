package routes

import (
	"github.com/bete7512/telegram-cms/handlers"
	"github.com/bete7512/telegram-cms/repositories"
	"github.com/bete7512/telegram-cms/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func userRoutes(v1 *gin.RouterGroup, db gorm.DB) {
	userService := services.NewUserService(repositories.NewDB(&db))
	userHandlers := handlers.NewUserHandlers(*userService)
	v1.GET("/users", userHandlers.GetAllUsers)
	v1.GET("/users/:id", userHandlers.GetUserByID)
	v1.POST("/users", userHandlers.CreateUser)
	v1.PUT("/users/:id", userHandlers.UpdateUser)
	v1.DELETE("/users/:id", userHandlers.DeleteUser)
}
