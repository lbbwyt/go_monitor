package oshandler

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestCollectErrorLog(t *testing.T) {
	go CollectErrorLog("./logfile.txt")
	for logMsg := range errChan {
		//两次消息发送的时间间隔必须大于5s，否则阿里会禁用账号10分钟。
		fmt.Println(logMsg)
	}
}

func TestPraseMsg(t *testing.T) {
	msg := "monitorlog_sys_999_" + "端口8001" + "绑定的服务异常"
	fmt.Println(strings.Split(msg, "_")[2])
	level, err := strconv.ParseInt(strings.Split(msg, "_")[2], 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(level)
}
