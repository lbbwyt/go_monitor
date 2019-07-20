package oshandler

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

var (
	errChan = make(chan string, 2000)
)

func CollectErrorLog(filename string) {
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		logs.Error("tail file err:", err)
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
