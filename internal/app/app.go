package app

import (
	"promotion/configs"
	"promotion/internal/controller"
	"promotion/internal/middleware"
	"promotion/internal/router"
	"promotion/internal/server"

	"github.com/gin-gonic/gin"
)

func Run(cfg *configs.Config) {
	infra := initInfra(cfg)
	engine := initServerDeps(cfg, infra)
	s := server.New(cfg, infra.log, engine)
	s.Start()
}

func initServerDeps(cfg *configs.Config, infra *infrastructure) *gin.Engine {
	authMiddlewares := middleware.New(cfg, infra.db.Firebase)
	modules := initModules(infra)
	controllers := initControllers(infra, modules)
	return router.NewEngine(cfg, infra.log, controllers, authMiddlewares)
}

func initControllers(infra *infrastructure, modules *Modules) *controller.Controllers {
	return &controller.Controllers{
		HealthCheck:  controller.NewHealthCheckController(),
		ReusableCode: controller.NewReusableCodeController(infra.log, modules.ReusableCode),
	}
}
