package handler

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_monitor/src/server/admin/config"
	"net/url"
	"testing"
)

func TestGetSmsAccessToken(t *testing.T) {
	config.InitConfig("F:/go/src/go_monitor/src/server/admin")

	token, err := GetSmsAccessToken()
	if err != nil {
		logs.Error("获取token失败" + err.Error())
	}
	fmt.Println(token)

}

func TestPushSmsMsg(t *testing.T) {
	config.InitConfig("F:/go/src/go_monitor/src/server/admin")
	PushSmsMsg("信息测试", "13631196027")
}

func TestToGb2132(t *testing.T) {
	fmt.Println(ToGb2132("你好"))

	fmt.Println(url.QueryUnescape(ToGb2132("你好")))
}
