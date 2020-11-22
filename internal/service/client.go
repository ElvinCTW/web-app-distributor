package service

import (
	"web-app-distributor/internal/dao"
	"web-app-distributor/pkg/logger"
)

var (
	log = logger.Get()
)

func Init() {
	dao.Init()
}
