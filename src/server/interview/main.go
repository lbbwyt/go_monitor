package main

import (
	log "github.com/astaxie/beego/logs"
	"go_monitor/src/server/interview/dynamic"
)

func main() {
	log.Info("Enter main")
	Init()

}

//初始化
func Init() {
	log.Info(">>>>开始初始化")
	log.Info(dynamic.RobotMN(4, 4))

}
