package handler

import (
	"chief_operation/src/server/admin/config"
	"github.com/astaxie/beego/logs"
	"time"
)

const (
	chan_dingding = "dingding"
	err_name_count = "count"
	count_interval = 60
)

type Message interface {
	Send(*Param) error
}

type Handle struct {
	mess    Message
	ChParam chan *Param
}

type Param struct {
	Org string  `json:"Org"` //单位名称 数字高新
	AppName string `json:"AppName"` //应用名称 monitor
	ModuleName string  `json:"ModuleName"`// 模块名称 智慧消防
	Code  string `json:"Code"` // 错误码
	Msg   string  `json:"Msg"`// 系统错误消息，较详细
	Error string  `json:"Error"`//自定义错误消息，
}

var handle *Handle

func InitService() {
	logs.Info("开始初始化钉钉服务")
	handle = new(Handle)
	handle.ChParam = make(chan *Param, 2000)
	handle.mess = NewDingDing(config.Conf.Dingding.Path)

	go Receive()
}

func Add(param *Param) {
	if config.Conf.Dingding.Send == 1 { //如果发送未开启，则不发送
		return
	}
	handle.ChParam <- param
}

func Receive() {
	for param := range handle.ChParam {
		//两次消息发送的时间间隔必须大于5s，否则阿里会禁用账号10分钟。
		time.Sleep(time.Second * 6)
		handle.mess.Send(param)
	}
}

func GetPeople() []string {
	arr := config.Conf.Dingding.People
	return arr
}
