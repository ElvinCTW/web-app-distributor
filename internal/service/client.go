package service

import (
	"web-app-distributor/config"
	"web-app-distributor/internal/dao"
)

var c = config.Get()

func Init() {
	dao.Init()
}
