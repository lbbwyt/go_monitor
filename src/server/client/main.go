package main

import (
	log "github.com/astaxie/beego/logs"
	"go_monitor/src/server/client/config"
	"go_monitor/src/server/client/dao"
	"go_monitor/src/server/client/oshandler"
)

func main() {
	log.Info("Enter main")
	Init()

}

//初始化
func Init() {
	log.Info(">>>>开始初始化")
	config.InitVipConfig()
	dao.InitMysql()

	oshandler.InitService()
	oshandler.InitLogService()
	//阻塞主线程
	var chanInt = make(chan int)
	<-chanInt

}
