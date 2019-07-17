package api

import (
	"chief_operation/src/server/admin/config"
	"chief_operation/src/server/admin/dao"
	"chief_operation/src/server/admin/handler"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"testing"
	"time"
)

func TestInitApi(t *testing.T) {
	config.InitConfig("F:/go/src/chief_operation/src/server/admin")
	dao.InitMysql();
	go InitApi("localhost:6800")
	time.Sleep(time.Second * 2)


	req := httplib.Post("http://localhost:6800/api/pushMsg")
	req.Header("Content-Type", "application/json")
	var param = new(handler.Param);
	param.Code= "402"
	param.Error = "test err msg"

	req.Body(param)
	res, err := req.Bytes()
	if err != nil {
		logs.Error("发送钉钉消息失败" + err.Error())
	}
	logs.Info("钉钉消息发送成功 %s", string(res))

	//阻塞主线程
	var chanInt = make(chan int)
	<-chanInt
}
