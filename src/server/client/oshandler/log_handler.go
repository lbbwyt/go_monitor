package oshandler

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"go_monitor/src/server/admin/entity"
	"go_monitor/src/server/client/config"
	"go_monitor/src/server/client/dao"
	emap "go_monitor/src/util/ExpiredMap"
	"go_monitor/src/util/common"
	"go_monitor/src/util/errors"
	"strconv"
	"strings"
	"time"
)

/*
tail源碼分析
https://www.cnblogs.com/zhaof/p/9663350.html
*/

var (
	errChan     = make(chan string, 2000)
	logCacheMap = emap.NewExpiredMap()
)

func InitLogService() {
	logs.Info("开始初始化日志监控服务")
	common.CheckGoPanic()
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

	if strings.Contains(msg, "monitorlog") {

		//消息解析

		res := PraseMsg(msg[strings.Index(msg, "monitorlog"):len(msg)])
		//持久化消息
		if res != nil {
			SaveMsg(res)
		}
	}
}

/**
* 解析日志消息为对应的持久化实体
 消息格式为：
 monitor_log_srvCode_level:msg
*/
func PraseMsg(msg string) *entity.AlarmMsg {
	var (
		res *entity.AlarmMsg
	)
	if len(strings.Split(msg, "_")) < 3 {
		return res
	}
	res = new(entity.AlarmMsg)
	res.AppName = "驾驶舱"
	res.ErrMsg = msg[11:len(msg)]
	res.IsVerfied = 0

	res.WarningId = strings.Split(msg, "_")[1]
	res.Level, _ = strconv.ParseInt(strings.Split(msg, "_")[2], 10, 64)
	res.StartTime = time.Now()
	res.CreateTime = time.Now()
	return res
}

func SaveMsg(msg *entity.AlarmMsg) error {
	if msg == nil {
		return errors.EntityIsNil()
	}
	errmsg := msg.ErrMsg
	logs.Info("**********************" + msg.ErrMsg + "**********************")

	index := 20
	if len(errmsg) < 20 {
		index = len(errmsg)
	}

	content := errmsg[0:index]
	//缓存msg的前面的数据，避免重复推送
	found, value := logCacheMap.Get(content)
	if found {
		logs.Info(value.(string) + " already alarm, ttl：")
		logs.Info(logCacheMap.TTL(content))
		return nil
	}
	logCacheMap.Set(content, content, int64(4*60*60))

	//写数据到数据库
	err := dao.MysqlCon.Save(msg).Error
	if err != nil {
		logs.Error("写入消息失败" + err.Error())
		return err
	}
	return nil

}

func PushDingLogMsg(errMsg string) error {
	url := config.Conf.Url
	logs.Info("api接口为：", url)
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
		logs.Error("发送钉钉消息失败,msg:" + errMsg + "err:" + err.Error())
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
