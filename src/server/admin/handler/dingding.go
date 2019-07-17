package handler

import (
    log "github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type DingDing struct {
	Path string
}

type DingContent struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewDingDing(path string) *DingDing {
	d := new(DingDing)
	d.Path = path
	return d
}

func (this *DingDing) Send(param *Param) error {
	data := this.SetParam(param)
	log.Info("开始发送钉钉消息" +  string(data))
	req := httplib.Post(this.Path)
	req.Header("Content-Type", "application/json")
	req.Body(data)
	res, err := req.Bytes()
	if err != nil {
		logs.Error("发送钉钉消息失败")
		return err
	}
	logs.Info("钉钉消息发送成功 %s", string(res))
	return nil
}

func (this *DingDing) SetParam(param *Param) []byte {
	content := new(DingContent)
	content.Msgtype = "text"
	content.Text.Content = fmt.Sprintf("异常单位：%v,异常应用：%v, 异常模块：%v,错误码为：%v, 错误提示为：%v, 具体信息为：%v",
		param.Org, param.AppName, param.ModuleName,
		param.Code, param.Error, param.Msg)

	content.At.AtMobiles = GetPeople()
	b, _ := json.Marshal(content)
	return b
}
