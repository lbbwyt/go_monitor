package oshandler

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"go_monitor/src/server/client/config"
	"go_monitor/src/util/common"
	"strings"
	"time"
)

/*
tail源碼分析
https://www.cnblogs.com/zhaof/p/9663350.html
*/

var (
	errChan = make(chan string, 2000)
)

func InitLogService() {
	logs.Info("开始初始化日志监控服务")
	var (
		path = config.Conf.Logpath
	)

	logs.Info("日志路径为：" + path)
	go CollectErrorLog(path)

	for {
		select {
		case msg := <-errChan:
			{
				HandleLogMsg(msg)
			}
		}
	}
}

func HandleLogMsg(msg string) {
	fmt.Println(msg)
	if strings.HasPrefix(msg, "monitor_log") {
		//推送异常消息
		logs.Info("************************" + msg + "*******************")
		PushDingLogMsg(msg)
	}
}

func PushDingLogMsg(errMsg string) error {
	url := config.Conf.Url
	org := config.Conf.Org
	content := new(common.Param)
	content.Org = org
	content.Msg = errMsg
	data, _ := json.Marshal(content)
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json")
	req.Body(data)
	res, err := req.Bytes()
	if err != nil {
		logs.Error("发送钉钉消息失败,msg:" + errMsg)
		return err
	}
	logs.Info("钉钉消息发送成功 %s", string(res))
	return nil
}

func CollectErrorLog(filename string) {
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		logs.Error("tail file err:", err.Error())
		return
	}

	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filenam:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		errChan <- msg.Text
	}
}
