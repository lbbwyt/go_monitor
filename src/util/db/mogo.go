package db

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	log "github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2"
	"src/gopkg.in/mgo.v2/bson"
	"strconv"
)

type Address struct {
	Id          int64  `bson:"_id"`
	GroupId     int64  `bson:"groupId"`
	AddressPath string `bson:"address"`
}

func Ccdb2() (*mgo.Database, *mgo.Session) {

	mgo_url := "mongodb://tarsa:tsps2018@192.168.8.10:27017/admin?authMechanism=SCRAM-SHA-1"

	session, err := mgo.Dial(mgo_url)

	session.SetMode(mgo.Monotonic, true)

	db := session.DB("iot") //数据库名称

	if err != nil {
		fmt.Println("------连接数据库失败------------")
		panic(err)
	}
	fmt.Println("------ConnectionDb-----2-------")
	return db, session
}

func PersonDocument() (*mgo.Collection, *mgo.Session) {
	db, session := Ccdb2()
	conn := db.C("iot_address_info")
	return conn, session
}

//groupID :
//团结社区：239424039298891777
//好景社区：239424424587657218
//新城社区：244828949174198273
//鹤州社区：242330343963803649

func GetBusiness() {
	fmt.Println("\n------------查询出一个集合--------------")
	var addresses []Address

	conn, session := PersonDocument()
	//conf.PersonDocument().Find(bson.M{"name": "aa"}).All(&person)
	// err := conn.Find(nil).All(&addresses)
	err := conn.Find(bson.M{"groupId": 242330343963803649}).All(&addresses)

	if err != nil {
		fmt.Println(err.Error())
	}
	for index, add := range addresses {
		fmt.Println("序号" + string(index) + add.AddressPath + "开始新增管理员")
		//AddManager(add)
		DelManager(add)
		//break;
	}
	session.Close() //关闭数据库
}

func AddManager(add Address) error {
	data := GetPostData(add)
	//调用物联网平台接口：
	err := PostData(data)
	if err != nil {
		log.Error("商户：" + add.AddressPath + "新增失败")
	}
	return nil
}

func DelManager(add Address) error {
	data := GetDelPostData(add)
	//调用物联网平台接口：
	err := PostDelData(data)
	if err != nil {
		log.Error("商户：" + add.AddressPath + "删除失败")
	}
	return nil
}

func GetDelPostData(address Address) []byte {
	manager := &Manager{
		UserId: "245552937802309634",
	}
	var managers = make([]Manager, 0)
	managers = append(managers, *manager)
	addressPostData := new(AddressPostData)
	addressPostData.AddressId = strconv.FormatInt(address.Id, 10)
	addressPostData.ManagerUsers = managers
	b, _ := json.Marshal(addressPostData)
	fmt.Println(string(b))
	return b
}

func GetPostData(address Address) []byte {
	manager := &Manager{
		UserId:    "251054877877358593",
		Mobile:    "15626920990",
		UserName:  "李勋强",
		Master:    "2",
		AlarmType: "1",
	}
	var managers = make([]Manager, 0)
	managers = append(managers, *manager)
	addressPostData := new(AddressPostData)
	addressPostData.AddressId = strconv.FormatInt(address.Id, 10)
	addressPostData.ManagerUsers = managers
	b, _ := json.Marshal(addressPostData)
	fmt.Println(string(b))
	return b
}

type AddressPostData struct {
	ManagerUsers []Manager `json:"managerUsers"`
	AddressId    string    `json:"addressId"`
}

type Manager struct {
	UserId    string `json:"userId"`
	Mobile    string `json:"mobile"`
	UserName  string `json:"userName"`
	Master    string `json:"master"`
	AlarmType string `json:"alarmType"`
}

func PostData(data []byte) error {
	log.Info("开始新增管理员" + string(data))
	path := "https://iot.chiefdata.net/iot/firectl/business/info/manager/add"
	cookie := "__sid=eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiIyMzY1NDMyODIwODc4MjEzMTMiLCJpcCI6IjE5Mi4xNjguOC44IiwiaXNzIjoidGFzIiwiaWF0IjoxNTY5NzQ1NTU0fQ.r-dQ8eqwLfaWqQECPpBOigFCcA_bT8oBAIspkF3scfDwBDGAoZd9j7efeeYcg05wbix5eQR5LROMeT80u_82Uw"
	req := httplib.Post(path)
	req.Header("Content-Type", "application/json")
	req.Header("Cookie", cookie)
	req.Body(data)
	res, err := req.Bytes()
	if err != nil {
		log.Error("新增管理员失败")
		return err
	}
	log.Info("新增管理员成功 %s", string(res))
	return nil
}

func PostDelData(data []byte) error {
	log.Info("开始删除管理员" + string(data))
	path := "https://iot.chiefdata.net/iot/firectl/business/info/manager/del"
	cookie := "__sid=eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiIyMzY1NDMyODIwODc4MjEzMTMiLCJpcCI6IjE5Mi4xNjguOC44IiwiaXNzIjoidGFzIiwiaWF0IjoxNTY5NzQ1NTU0fQ.r-dQ8eqwLfaWqQECPpBOigFCcA_bT8oBAIspkF3scfDwBDGAoZd9j7efeeYcg05wbix5eQR5LROMeT80u_82Uw"
	req := httplib.Post(path)
	req.Header("Content-Type", "application/json")
	req.Header("Cookie", cookie)
	req.Body(data)
	res, err := req.Bytes()
	if err != nil {
		log.Error("新增管理员失败")
		return err
	}
	log.Info("删除管理员成功 %s", string(res))
	return nil
}
