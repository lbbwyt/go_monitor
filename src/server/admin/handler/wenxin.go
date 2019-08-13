package handler

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"go_monitor/src/server/admin/config"
	emap "go_monitor/src/util/ExpiredMap"
)

var (
	ACCESS_TOKEN_URL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	PUSH_URL         = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
)

type WxResData struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

var (
	ExpiredMap = emap.NewExpiredMap()
)

func GetAccessToken() (string, error) {
	found, token := ExpiredMap.Get("wx_access_token")
	if found {
		return token.(string), nil
	}

	corpId := config.Conf.Weixin.Corpid
	corpSecret := config.Conf.Weixin.Corpsecret
	url := fmt.Sprintf(ACCESS_TOKEN_URL, corpId, corpSecret)
	req := httplib.Get(url)

	res, err := req.Bytes()
	if err != nil {
		logs.Info("获取accessToken失败")
		return "", err
	}
	v := new(WxResData)
	err = json.Unmarshal(res, v)
	if err != nil {
		logs.Error("解析accesstoken返回数据失败")
		return "", err
	}
	logs.Info("获取token为" + v.Access_token)
	ExpiredMap.Set("wx_access_token", v.Access_token, int64(v.Expires_in))
	return v.Access_token, nil
}

type WxContent struct {
	Touser  string `json:"touser"`
	ToParty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe            int `json:"safe"`
	Enable_id_trans int `json:"enable_id_trans"`
}

func PushWxMsg(content string) error {
	var (
		token, _ = GetAccessToken()
	)
	url := fmt.Sprintf(PUSH_URL, token)
	msg := new(WxContent)
	msg.Touser = "@all"
	msg.Msgtype = "text"
	msg.Agentid = config.Conf.Weixin.Agentid
	msg.Text.Content = content
	b, _ := json.Marshal(msg)
	logs.Info("开始发送微信消息")
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json")
	req.Body(b)
	res, err := req.Bytes()
	if err != nil {
		logs.Error("发送微信消息失败")
		return err
	}
	logs.Info("微信消息发送成功 %s", string(res))
	return nil
}
