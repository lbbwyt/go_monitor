package handler

import (
	"go_monitor/src/server/admin/config"
	"testing"
)

func TestInitService(t *testing.T) {
	config.InitConfig("F:/go/src/chief_operation/src/server/admin")
	InitService()
	SendMsg()

	//阻塞主线程
	var chanInt = make(chan int)
	<-chanInt
}

func SendMsg() {
	for i := 0; i < 10; i++ {
		var param = new(Param)
		param.Msg = "服务器异常"
		param.Code = "502"
		param.Error = "test"
		Add(param)
	}
}
