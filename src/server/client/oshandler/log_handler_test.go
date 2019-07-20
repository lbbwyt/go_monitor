package oshandler

import (
	"fmt"
	"testing"
)

func TestCollectErrorLog(t *testing.T) {
	go CollectErrorLog("./logfile.txt")
	for logMsg := range errChan {
		//两次消息发送的时间间隔必须大于5s，否则阿里会禁用账号10分钟。
		fmt.Println(logMsg)
	}
}
