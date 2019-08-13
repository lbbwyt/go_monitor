package main

import (
	log "github.com/astaxie/beego/logs"
	"go_monitor/src/server/admin/api"
	"go_monitor/src/server/admin/config"
	"go_monitor/src/server/admin/dao"
	"go_monitor/src/server/admin/handler"
)

func main() {
	log.Info("Enter main")
	Init()

}

//初始化
func Init() {
	log.Info(">>>>开始初始化")
	//debug.InitDebug("localhost:6060")
	config.InitVipConfig()
	dao.InitMysql()
	handler.InitService()
	api.InitApi(":6800")

}
