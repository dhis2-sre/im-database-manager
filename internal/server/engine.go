package server

import (
	"net/http/pprof"

	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/internal/middleware"
	"github.com/dhis2-sre/im-database-manager/pkg/database"
	"github.com/dhis2-sre/im-database-manager/pkg/health"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redocMiddleware "github.com/go-openapi/runtime/middleware"
)

func GetEngine(basePath string, dbHandler database.Handler, authMiddleware handler.AuthenticationMiddleware) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("authorization")
	r.Use(cors.New(corsConfig))

	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(basePath)

	redoc(router, basePath)

	router.GET("/health", health.Health)

	router.GET("/databases/external/:uuid", dbHandler.ExternalDownload)

	tokenAuthenticationRouter := router.Group("")
	tokenAuthenticationRouter.Use(authMiddleware.TokenAuthentication)
	tokenAuthenticationRouter.POST("/databases", dbHandler.Upload)
	tokenAuthenticationRouter.POST("/databases/:id/copy", dbHandler.Copy)
	tokenAuthenticationRouter.GET("/databases/:id/download", dbHandler.Download)
	tokenAuthenticationRouter.GET("/databases", dbHandler.List)
	tokenAuthenticationRouter.GET("/databases/:id", dbHandler.FindById)
	tokenAuthenticationRouter.PUT("/databases/:id", dbHandler.Update)
	tokenAuthenticationRouter.DELETE("/databases/:id", dbHandler.Delete)
	tokenAuthenticationRouter.POST("/databases/:id/lock", dbHandler.Lock)
	tokenAuthenticationRouter.DELETE("/databases/:id/unlock", dbHandler.Unlock)
	tokenAuthenticationRouter.POST("/databases/save-as/:instanceId", dbHandler.SaveAs)
	tokenAuthenticationRouter.POST("/databases/:id/external", dbHandler.CreateExternalDownload)

	pfRouter := router.Group("/debug/pprof")
	pfRouter.GET("/", gin.WrapF(pprof.Index))
	pfRouter.GET("/cmdline", gin.WrapF(pprof.Cmdline))
	pfRouter.GET("/profile", gin.WrapF(pprof.Profile))
	// TODO add POST /symbol ?
	pfRouter.GET("/symbol", gin.WrapF(pprof.Symbol))
	pfRouter.GET("/trace", gin.WrapF(pprof.Trace))
	pfRouter.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
	pfRouter.GET("/heap", gin.WrapH(pprof.Handler("heap")))
	pfRouter.GET("/block", gin.WrapH(pprof.Handler("block")))
	pfRouter.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
	pfRouter.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))

	return r
}

func redoc(router *gin.RouterGroup, basePath string) {
	router.StaticFile("/swagger.yaml", "./swagger/swagger.yaml")

	redocOpts := redocMiddleware.RedocOpts{
		BasePath: basePath,
		SpecURL:  "./swagger.yaml",
	}
	router.GET("/docs", func(c *gin.Context) {
		redocHandler := redocMiddleware.Redoc(redocOpts, nil)
		redocHandler.ServeHTTP(c.Writer, c.Request)
	})
}
