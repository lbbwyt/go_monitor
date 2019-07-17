package main

import (
	"chief_operation/src/server/client/config"
	"chief_operation/src/server/client/oshandler"
	log "github.com/astaxie/beego/logs"
)

func main() {
	log.Info("Enter main")
	Init();

}

//初始化
func Init() {
	log.Info(">>>>开始初始化")
    config.InitConfig("")
	oshandler.InitService()
	//阻塞主线程
	var chanInt = make(chan int)
	<-chanInt

}
