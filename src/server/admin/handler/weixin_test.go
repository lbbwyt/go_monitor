package handler

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"testing"
)
import "go_monitor/src/server/admin/config"

func TestGetAccessToken(t *testing.T) {
	config.InitConfig("F:/go/src/go_monitor/src/server/admin")

	token, err := GetAccessToken()
	if err != nil {
		logs.Error("获取token失败" + err.Error())
	}
	fmt.Println(token)

}

func TestPushWxMsg(t *testing.T) {
	config.InitConfig("F:/go/src/go_monitor/src/server/admin")
	PushWxMsg("驾驶舱异常监控微信推送测试消息")
}
