package db

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"go_monitor/src/util/common"
	"gopkg.in/mgo.v2"
	"src/gopkg.in/mgo.v2/bson"
)

func DeviceInfoDocument() (*mgo.Collection, *mgo.Session) {
	db, session := Ccdb2()
	conn := db.C("iot_device_info")
	return conn, session
}

func AlarmInfoDocument() (*mgo.Collection, *mgo.Session) {
	db, session := Ccdb2()
	conn := db.C("busi_alarm_info")
	return conn, session
}

func BusinessInfoDocument() (*mgo.Collection, *mgo.Session) {
	db, session := Ccdb2()
	conn := db.C("busi_business_info")
	return conn, session
}

type Device struct {
	Id                int64  `bson:"_id"`
	Address           string `bson:"address"`
	TypeName          string `bson:"typeName"`
	DeviceName        string `bson:"deviceName"`
	AddressId         int64  `bson:"addressId"`
	FullGroupPathName string `bson:"fullGroupPathName"`
	Latitude          string `bson:"latitude"`
	Longitude         string `bson:"longitude"`
	Pname             string `bson:"pname"`
	CreateDate        string `bson:"createDate"`
	ImeiNo            string `bson:"imeiNo"`
}

type Business struct {
	Id            int64  `bson:"_id"`
	BusinessName  string `bson:"businessName"`
	ContactPerson string `bson:"contactPerson"`
	AddressId     int64  `bson:"addressId"`
	CreateDate    string `bson:"createDate"`
	ContactMobile int64  `bson:"contactMobile"`
}

type AlarmInfo struct {
	Id            int64  `bson:"_id"`
	DeviceId      int64  `bson:"deviceId"`
	FullAddress   string `bson:"fullAddress"`
	AddressId     int64  `bson:"addressId"`
	CreateDate    int64  `bson:"createDate"`
	DealInfo      string `bson:"dealInfo"`
	AlarmInfo     string `bson:"alarmInfo"`
	ImeiNo        string `bson:"imeiNo"`
	BusinessName  string
	ContactPerson string
	ContactMobile int64
}

type ExportData struct {
	Id                int64  `bson:"_id"`
	ImeiNo            string `bson:"imeiNo"`
	Address           string `bson:"address"`
	TypeName          string `bson:"typeName"`
	DeviceName        string `bson:"deviceName"`
	AddressId         int64  `bson:"addressId"`
	FullGroupPathName string `bson:"fullGroupPathName"`
	Latitude          string `bson:"latitude"`
	Longitude         string `bson:"longitude"`
	Pname             string `bson:"pname"`
	CreateDate        string `bson:"createDate"`
	BusinessName      string `bson:"businessName"`
	ContactPerson     string `bson:"contactPerson"`
	ContactMobile     int64  `bson:"contactMobile"`
}

/**
 * 获取20190912之后新增的设备
 */
func GetDeviceInfo() []Device {
	fmt.Println("\n------------查询出一个集合--------------")
	var devices []Device
	conn, session := DeviceInfoDocument()
	defer session.Close() //关闭数据库

	err := conn.Find(bson.M{"createDate": bson.M{"$gte": 20190912}}).All(&devices)

	if err != nil {
		fmt.Println(err.Error())
	}
	return devices
}

/**
 * 获取20190912之后新增的设备
 */
func GetBusinessInfo() []Business {
	fmt.Println("\n------------查询出一个集合--------------")
	var businesses []Business
	conn, session := BusinessInfoDocument()
	defer session.Close() //关闭数据库

	//err := conn.Find(bson.M{"fullGroupPathName": bson.M{"$regex": bson. RegEx:{Pattern:"/a/", Options: "im"}} }).All(&businesses)
	//err := conn.Find(bson.M{"fullGroupPathName": bson.M{"$regex": bson.RegEx{Pattern:"/斗门/", Options: "im"}}}).All(&businesses)
	err := conn.Find(bson.M{"fullGroupPathName": bson.M{"$regex": "斗门"}}).All(&businesses)
	if err != nil {
		fmt.Println(err.Error())
	}
	return businesses
}

func GetAlarmInfo() []AlarmInfo {
	fmt.Println("\n------------查询出一个集合--------------")
	var alarmInfos []AlarmInfo
	conn, session := AlarmInfoDocument()
	defer session.Close() //关闭数据库

	//err := conn.Find(bson.M{"fullGroupPathName": bson.M{"$regex": bson. RegEx:{Pattern:"/a/", Options: "im"}} }).All(&businesses)
	//err := conn.Find(bson.M{"fullGroupPathName": bson.M{"$regex": bson.RegEx{Pattern:"/斗门/", Options: "im"}}}).All(&businesses)
	err := conn.Find(bson.M{"createDate": bson.M{"$gte": 20191001}, "fullAddress": bson.M{"$regex": "斗门"}}).All(&alarmInfos)
	if err != nil {
		fmt.Println(err.Error())
	}
	return alarmInfos
}

func GetAlarmExportData() {
	var (
		businesses  []Business
		alarmInfos  []AlarmInfo
		exportDatas []AlarmInfo
	)
	alarmInfos = GetAlarmInfo()
	businesses = GetBusinessInfo()
	var addressId2Business = make(map[int64]Business, 0)

	for _, item := range businesses {
		addressId2Business[item.AddressId] = item
	}

	for _, item := range alarmInfos {
		item.BusinessName = addressId2Business[item.AddressId].BusinessName
		item.ContactPerson = addressId2Business[item.AddressId].ContactPerson
		item.ContactMobile = addressId2Business[item.AddressId].ContactMobile

		exportDatas = append(exportDatas, item)
	}

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("设备信息")
	for i, data := range exportDatas {
		str := fmt.Sprintf("%s,%d", data.BusinessName, data.ContactMobile)
		fmt.Println(str)
		row := sheet.AddRow()
		row.SetHeightCM(1) //设置每行的高度
		cell := row.AddCell()
		cell.Value = common.Int2String(i + 1)
		//id
		cell = row.AddCell()
		cell.Value = data.ImeiNo
		//型号

		cell = row.AddCell()
		cell.Value = data.AlarmInfo
		//类型
		cell = row.AddCell()
		cell.Value = data.DealInfo
		//设备名称
		cell = row.AddCell()
		cell.Value = data.FullAddress
		//地址1
		cell = row.AddCell()
		cell.Value = data.BusinessName
		//地址2
		cell = row.AddCell()
		cell.Value = data.ContactPerson
		//商户名称
		cell = row.AddCell()
		cell.Value = common.Int642String(data.CreateDate)

	}
	err := file.Save("file.xlsx")
	if err != nil {
		panic(err)
	}

}

func GetExportData() {
	var (
		businesses []Business
		devices    []Device
	)
	devices = GetDeviceInfo()
	businesses = GetBusinessInfo()
	var addressId2Business = make(map[int64]Business, 0)

	for _, item := range businesses {
		addressId2Business[item.AddressId] = item
	}

	var expertDatas = make([]ExportData, 0)

	for _, device := range devices {
		_, ok := addressId2Business[device.AddressId]
		if !ok {
			continue
		}
		var expert = new(ExportData)
		expert.AddressId = device.AddressId
		expert.Address = device.Address
		expert.CreateDate = device.CreateDate
		expert.Id = device.Id
		expert.DeviceName = device.DeviceName
		expert.BusinessName = addressId2Business[device.AddressId].BusinessName
		expert.ContactPerson = addressId2Business[device.AddressId].ContactPerson
		expert.ContactMobile = addressId2Business[device.AddressId].ContactMobile
		expert.FullGroupPathName = device.FullGroupPathName
		expert.Pname = device.Pname
		expert.ImeiNo = device.ImeiNo
		expertDatas = append(expertDatas, *expert)
	}

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("设备信息")
	for i, data := range expertDatas {
		str := fmt.Sprintf("%s,%d", data.BusinessName, data.ContactMobile)
		fmt.Println(str)
		row := sheet.AddRow()
		row.SetHeightCM(1) //设置每行的高度
		cell := row.AddCell()
		cell.Value = common.Int2String(i + 1)
		//id
		cell = row.AddCell()
		cell.Value = data.ImeiNo
		//型号

		cell = row.AddCell()
		cell.Value = data.Pname
		//类型
		cell = row.AddCell()
		cell.Value = data.TypeName
		//设备名称
		cell = row.AddCell()
		cell.Value = data.DeviceName
		//地址1
		cell = row.AddCell()
		cell.Value = data.FullGroupPathName
		//地址2
		cell = row.AddCell()
		cell.Value = data.Address
		//商户名称
		cell = row.AddCell()
		cell.Value = data.BusinessName
		//商户联系人
		cell = row.AddCell()
		cell.Value = data.ContactPerson

		//商户电话
		cell = row.AddCell()
		cell.Value = common.Int642String(data.ContactMobile)

	}
	err := file.Save("file.xlsx")
	if err != nil {
		panic(err)
	}

}
