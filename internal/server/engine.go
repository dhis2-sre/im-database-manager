package server

import (
	"github.com/dhis2-sre/im-database-manager/internal/di"
	"github.com/dhis2-sre/im-database-manager/internal/middleware"
	"github.com/dhis2-sre/im-database-manager/pkg/health"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redocMiddleware "github.com/go-openapi/runtime/middleware"
)

func GetEngine(environment di.Environment) *gin.Engine {
	basePath := environment.Config.BasePath

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(basePath)

	redoc(router, basePath)

	//	router.GET("/health", health.Health)

	tokenAuthenticationRouter := router.Group("")
	tokenAuthenticationRouter.Use(environment.AuthenticationMiddleware.TokenAuthentication)
	tokenAuthenticationRouter.GET("/health", health.Health)

	return r
}

func redoc(router *gin.RouterGroup, basePath string) {
	router.StaticFile("/swagger.yaml", "./swagger/swagger.yaml")

	redocOpts := redocMiddleware.RedocOpts{
		BasePath: basePath,
		SpecURL:  basePath + "/swagger.yaml",
	}
	router.GET("/docs", func(c *gin.Context) {
		redocHandler := redocMiddleware.Redoc(redocOpts, nil)
		redocHandler.ServeHTTP(c.Writer, c.Request)
	})
}
