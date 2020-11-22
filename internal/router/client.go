package router

import (
	"github.com/gin-gonic/gin"
	"web-app-distributor/internal/service"
	"web-app-distributor/pkg/logger"
)

var log = logger.Get()

func Init(r *gin.Engine) {
	service.Init()
	Article(r)
}
