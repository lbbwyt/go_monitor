package oshandler

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"go_monitor/src/server/client/config"
	"go_monitor/src/util/common"
	"os/exec"
	"time"
)

func InitService() {
	logs.Info("开始初始化进程服务")
	ports := config.Conf.Ports
	for _, port := range ports {
		go StartMonitorPort(port)
	}
}

/**
定时检测进程存活常规下每隔一分钟检测一次。
*/
func StartMonitorPort(port string) {
	tickerMonitor := time.NewTicker(60 * time.Second)
	go func(t time.Ticker) {
		for {
			<-t.C
			logs.Info("循环检活,端口号为"+port, time.Now().Format("2006-01-02 15:04:05"))
			execPipleCmd(port)
		}
	}(*tickerMonitor)
}

func execPipleCmd(port string) {
	cmds := []*exec.Cmd{
		exec.Command("netstat", "-tlnp"),
		exec.Command("grep", port),
	}

	outstr, err := ExecPipeLine(cmds...)
	if err != nil {
		msg := "monitorlog_sys_999_" + "端口" + port + "绑定的服务异常"

		res := PraseMsg(msg)
		//持久化消息
		if res != nil {
			SaveMsg(res)
		}
	}

	logs.Info("linux 命令执行结果为" + outstr)
}

func PushDingMsg(errMsg string, port string) error {
	url := config.Conf.Url
	org := config.Conf.Org
	content := new(common.Param)
	content.Org = org
	content.Msg = "端口号为（" + port + "）的服务进程已死亡,请及时处理" + errMsg
	data, _ := json.Marshal(content)
	req := httplib.Post(url)
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
