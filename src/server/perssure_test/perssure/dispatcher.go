package perssure

import (
	log "github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type Dispatcher struct {
	Num    int64
	Robots []*Robot
}

/**
实例化分发器，包含一定数量的robot
*/
func NewDispatcher(num int64, clientId string) *Dispatcher {
	dispatcher := new(Dispatcher)
	dispatcher.Num = num
	robots := make([]*Robot, num)
	for index, r := range robots {
		r = NewRobot()
		if clientId != "" {
			r.ClientId = clientId
		} else {
			r.ClientId = strconv.Itoa(index + 1)
		}
		robots[index] = r
	}
	dispatcher.Robots = robots
	return dispatcher
}

/**
连接失败channel
*/
var ConnectFailChan = make(chan string)

/**
成功完成任务channel
*/
var SuccessChan = make(chan string)

/**
 * 发布消息
 */
func Start() {
	log.Info("开始执行压测")
	dispatcher := NewDispatcher(3000, "")

	for _, r := range dispatcher.Robots {
		go StartPublish(r)
		//
		time.Sleep(time.Millisecond * 10)
	}
}

func StartPublish(robot *Robot) {
	_, err := robot.Connect()
	if err != nil {
		return
	}

	log.Info("设备" + robot.ClientId + "开始发布消息")

	for j := 0; j < 100; j++ {
		for i := 0; i < 2; i++ {
			robot.Public()
		}
		time.Sleep(time.Second * 1)
	}
	log.Info("设备" + robot.ClientId + "结束发布消息")
	if robot.Client.IsConnected() {
		robot.DisConnect()
	} else {
		log.Error("异常，设备连接自动断开" + robot.ClientId)
	}
}
