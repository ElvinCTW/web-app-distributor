package router

import (
	"github.com/gin-gonic/gin"
	"web-app-distributor/config"
	"web-app-distributor/internal/service"
	"web-app-distributor/pkg/logger"
)

var (
	log         = logger.Get()
	apiVersion  = config.Get().API_VERSION
	appPrefix   = "/app/" + apiVersion
	adminPrefix = "/admin/" + apiVersion
)

func Init(r *gin.Engine) {
	service.Init()
	Article(r)
	Restaurant(r)
}
