package router

import (
	"promotion/configs"
	"promotion/internal/controller"
	"promotion/internal/middleware"
	"promotion/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewEngine(
	cfg *configs.Config,
	log *logger.Logger,
	controllers *controller.Controllers,
	authMiddlewares *middleware.AuthMiddlewares,
) *gin.Engine {
	if cfg.Server.Env == configs.ServerEnvProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gin.ForceConsoleColor()
	gin.DebugPrintRouteFunc = logger.DebugOutputLogger(log)

	engine := gin.New()
	attachMiddleware(engine, log)

	registerRoutes(engine, controllers, authMiddlewares)
	return engine
}

func attachMiddleware(engine *gin.Engine, log *logger.Logger) {
	engine.Use(middleware.ErrorHandler(log))
	engine.Use(middleware.LoggerMiddleware(log))
	engine.Use(middleware.RecoveryMiddleware(log))
}

func registerRoutes(
	engine *gin.Engine,
	controllers *controller.Controllers,
	authMiddlewares *middleware.AuthMiddlewares,
) {
	root := engine.Group("promotion")
	initHealthCheckRouter(root, controllers.HealthCheck)
	initReusableCodeRouter(root, controllers.ReusableCode, authMiddlewares.InternalAuth)
}
