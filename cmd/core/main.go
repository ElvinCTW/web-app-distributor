package main

import (
	"github.com/gin-gonic/gin"
	"web-app-distributor/config"
	"web-app-distributor/internal/router"
	"web-app-distributor/pkg/logger"
)

func main() {
	logger.Init(config.Get().LogLevel)
	r := gin.Default()
	router.Init(r)
	r.Run(":8080")
}
