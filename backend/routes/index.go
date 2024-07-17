package routes

import (
	"github.com/bete7512/telegram-cms/docs"
	"github.com/bete7512/telegram-cms/routes/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	routes := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	routes.Use(middleware.Prometheus())
	routes.Use(middleware.Logger())
	routes.Use(middleware.VerifyJwt())
	v1 := routes.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Ping)
		}
	}
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.Run(":8084")

	return routes
}

// PingExample godoc
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/example/helloworld [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


// TODO: add addPrometheusMetrics middleware
// TODO: add addLogger middleware
// TODO: add addRecovery middleware