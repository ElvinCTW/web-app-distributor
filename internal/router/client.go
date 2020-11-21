package router

import (
	"github.com/gin-gonic/gin"
	"web-app-distributor/internal/service"
)

func Set(r *gin.Engine) {
	service.Init()

}
