package routes

import (
	"github.com/bete7512/telegram-cms/handlers"
	"github.com/bete7512/telegram-cms/repositories"
	"github.com/bete7512/telegram-cms/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func authRoutes(v1 *gin.RouterGroup, db gorm.DB) {
	userService := services.NewUserService(repositories.NewDB(&db))
	authHandlers := handlers.NewAuthenticationHandlers(*userService)
	v1.POST("/signup", authHandlers.SignUp)
	v1.POST("/login", authHandlers.Login)
	v1.POST("/forget-password", authHandlers.ForgetPassword)
	v1.POST("/reset-password", authHandlers.ResetPassword)
	v1.POST("/change-password", authHandlers.ChangePassword)
	v1.POST("/verify-email", authHandlers.VerifyEmail)
}
