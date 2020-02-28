package perssure

import "fmt"
import MQTT "github.com/eclipse/paho.mqtt.golang"
import log "github.com/astaxie/beego/logs"

type Robot struct {
	ClientId   string
	ClientName string
	ProductId  string
	Client     MQTT.Client
}

/**
 * 构造函数
 */

func NewRobot() *Robot {
	d := new(Robot)
	d.ProductId = "test-fhsj-2.0"
	return d
}

/**
定义默认的消息处理器
*/
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Info(fmt.Printf("TOPIC: %s\n", msg.Topic()))
	log.Info(fmt.Printf("MSG: %s\n", msg.Payload()))
}

//{"devid":"866971030771930","pid":"H388N","pname":"H388N","cid":34,"aid":1,"a_name":"","bid":2,"b_name":"","lid":5,"l_name":"aaa","time":"2019-09-16 16:50:42","alarm_type":1,"alarm_type_name":"test","event_id":32,"event_count":1,"device_type":1,"comm_type":2,"first_alarm_time":"2019-09-16 16:50:42","last_alarm_time":"2019-09-16 16:50:42"}

/**
 * 发布消息
 */
func (this *Robot) Public() error {
	//log.Info("客户端开始发布消息，客户端Id" + this.ClientId)
	var topic = "/chiefdata/push/fire_alarm/department/1/area/2/dev/3"
	var text = `{"devid":"` + this.ClientId + `","pid":"H388N","pname":"H388N","cid":34,"aid":1,"a_name":"","bid":2,"b_name":"eeee","lid":5,"l_name":"aaa","time":"2019-09-16 16:50:42","alarm_type":1,"alarm_type_name":"test","event_id":32,"event_count":1,"device_type":1,"comm_type":2,"first_alarm_time":"2019-09-16 16:50:42","last_alarm_time":"2019-09-16 16:50:42"}`
	token := this.Client.Publish(topic, 0, false, text)
	token.Wait()
	return nil
}

/**
 * 连接MQTT服务
 */
func (this *Robot) Connect() (MQTT.Client, error) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.8.5:1883")
	opts.SetClientID(this.ClientId)
	opts.SetDefaultPublishHandler(f)
	opts.SetUsername("test")
	opts.SetPassword("test")
	//不设置端开重连
	//opts.SetAutoReconnect(true)

	opts.OnConnectionLost = func(client MQTT.Client, e error) {
		log.Error("连接断开:" + this.ClientId)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		//log.Error("客户端连接失败！" + this.ClientId)
		//var ConnectFailChan = make(chan string)
		//panic(token.Error())
		ConnectFailChan <- this.ClientId + ":" + token.Error().Error()
		return nil, token.Error()

	}
	this.Client = c
	return c, nil
}

/**
 * 断开MQTT服务连接
 */
func (this *Robot) DisConnect() {
	log.Info("客户端开始断开连接，客户端Id" + this.ClientId)
	this.Client.Disconnect(250)
	SuccessChan <- this.ClientId
}
