package main

import (
	"chief_operation/src/server/admin/api"
	"chief_operation/src/server/admin/config"
	"chief_operation/src/server/admin/dao"
	"chief_operation/src/server/admin/handler"
	log "github.com/astaxie/beego/logs"
)

func main() {
	log.Info("Enter main")
	Init();

}

//初始化
func Init() {
	log.Info(">>>>开始初始化")
	//debug.InitDebug("localhost:6060")
    config.InitConfig("")
	dao.InitMysql()
	handler.InitService()
	api.InitApi(":6800")

}
