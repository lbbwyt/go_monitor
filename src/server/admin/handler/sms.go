package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
	"go_monitor/src/server/admin/config"
	"gopkg.in/iconv.v1"
	"net/url"
)

var (
	SMS_ACCESS_TOKEN_URL = "%s/authtoken/gettoken"
	SMS_PUSH_URL         = "%s/sms/batchsend"
)

type SmsTokenParm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SmsPushMsgParm struct {
	Mobile  string `json:"mobile"`
	Content string `json:"content"`
	Svrtype string `json:"svrtype"`
	Exno    string `json:"exno"`
	Custid  string `json:"custid"`
	Exdata  string `json:"exdata"`
}

type SmsResData struct {
	Result      string `json:"result"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Token       string `json:"token"`
}

func GetSmsAccessToken() (string, error) {
	//found, token := ExpiredMap.Get("sms_access_token")
	//if found {
	//	return token.(string), nil
	//}
	host := config.Conf.Sms.Host
	url := fmt.Sprintf(SMS_ACCESS_TOKEN_URL, host)

	param := &SmsTokenParm{
		"admin",
		Md5V("P@ssw0rd520"),
	}
	b, _ := json.Marshal(param)
	logs.Info("获取sms短信token参数为：" + string(b))
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json")
	req.Body(b)
	res, err := req.Bytes()
	if err != nil {
		logs.Error("获取sms短信token失败" + err.Error())
		return "", err

	}
	logs.Info("获取sms短信token成功 %s", string(res))
	v := new(SmsResData)
	err = json.Unmarshal(res, v)
	if err != nil {
		logs.Error("解析accesstoken返回数据失败")
		return "", err
	}
	logs.Info("获取token为" + v.Token)
	return v.Token, nil
}

func PushSmsMsg(content string, phone string) error {
	token, err := GetSmsAccessToken()
	if err != nil {
		logs.Error("获取token失败！")
		return err
	}
	host := config.Conf.Sms.Host
	url := fmt.Sprintf(SMS_PUSH_URL, host)

	param := &SmsPushMsgParm{
		Mobile:  phone,
		Content: ToGb2132(content),
		Custid:  uuid.New().String(),
	}
	b, _ := json.Marshal(param)
	logs.Info("参数为：" + string(b))
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json")
	req.Header("Authorization", "Bearer "+token)
	req.Body(b)
	res, err := req.Bytes()
	if err != nil {
		logs.Error("发送sms消息失败")
		return err
	}
	logs.Info("sms发送成功 %s", string(res))
	return nil
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//gopkg.in/iconv.v1 包需要gcc的支持，windows下可按下述方法安装
//Win10配置：
//1.打开Powershell，win10自带，win7版本需要去微软官方下载补丁，是一个类似于Python pip的包管理装置，并需要以管理员身份运行：
//
//2.设置Get-ExecutionPolicy可用，PowerShell中输入：
//set-ExecutionPolicy RemoteSigned
//
//3.安装Chocolatey，这是一个第三方的包管理器，官方网址：https://chocolatey.org/
//iwr https://chocolatey.org/install.ps1 -UseBasicParsing | iex
//
//4：安装mingw
//choco install mingw
//
//上述过程很慢，请经常失败，需耐心等待。
//
//安装下面的包
//go get gopkg.in/iconv.v1
//
//
//类似于与java的java.net.URLEncoder.encode(msg,"gb2312")
func ToGb2132(str string) string {
	cd, err := iconv.Open("gb2312", "utf-8") // convert utf-8 to gb2312
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return str
	}
	defer cd.Close()

	gbk := cd.ConvString(str)

	return url.QueryEscape(gbk)
}
