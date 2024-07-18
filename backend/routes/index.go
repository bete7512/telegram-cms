package routes

import (
	"github.com/bete7512/telegram-cms/docs"
	"github.com/bete7512/telegram-cms/routes/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Router(db gorm.DB) *gin.Engine {

	routes := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	routes.Use(middleware.Prometheus())
	routes.Use(middleware.Logger())
	routes.Use(middleware.VerifyJwt())
	v1 := routes.Group("/api/v1")
	userRoutes(v1, db)
	authRoutes(v1, db)
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.Run(":8084")

	return routes
}
