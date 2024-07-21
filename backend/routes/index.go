package routes

import (
	"fmt"

	"github.com/bete7512/telegram-cms/config"
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
	v1 := routes.Group("/api/v1")
	authRoutes(v1, db)
	userRoutes(v1, db)
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.Run(fmt.Sprintf(":%s", config.PORT))

	return routes
}
