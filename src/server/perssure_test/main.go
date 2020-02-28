package main

import (
	log "github.com/astaxie/beego/logs"
	"go_monitor/src/server/perssure_test/perssure"
	"strconv"
)

func main() {
	log.Info("Enter main")
	Init()

}

//初始化
func Init() {
	log.Info(">>>>开始初始化")
	perssure.Start()
	var sucCount = 0

	////阻塞主线程
	for {

		select {
		case clientId := <-perssure.ConnectFailChan:
			log.Error("设备" + clientId + "连接失败")
		case <-perssure.SuccessChan:
			sucCount = sucCount + 1
			log.Info("完结数：" + strconv.Itoa(sucCount))
		}

	}
}
