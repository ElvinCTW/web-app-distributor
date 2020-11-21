package cmd

import (
	"github.com/gin-gonic/gin"
	"web-app-distributor/config"
	"web-app-distributor/internal/router"
	"web-app-distributor/pkg/logger"
)

func main() {
	logger.Init(config.Get().LogLevel)
	r := gin.Default()
	router.Set(r)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Run(":8080")
}
