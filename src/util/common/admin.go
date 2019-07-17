package common

import (
	"fmt"
	"strconv"
	"strings"
	"net"
	"os"
)

const (
	Admin_Upload_Path       string = "/srv/upload/admin/" //管理后台上传图片路径
	Short_Admin_Upload_Path string = "/upload/admin/"                    //管理后台上传图片短路径。用于回写数据库
)

func Typeof(v interface{}) string {
	t := fmt.Sprintf("%T", v)
	return t
}

func Interface2String(v interface{}) string {
	switch Typeof(v) {
	case "float64":
		value := int(v.(float64))
		return strconv.Itoa(value)
	case "string":
		return v.(string)
	case "int":
		return strconv.Itoa(v.(int))

	case "int32":
		return strconv.Itoa(int(v.(int32)))
	case "int64":
		return strconv.Itoa(int(v.(int64)))
	default:
		return ""
	}
}

//版本号字符串转数字
func Version2Num(version string) int64 {
	if version == "" {
		return 0
	}
	strs := strings.Split(version, ".")
	var versionNum int64
	len := len(strs) - 1
	for _, item := range strs {
		strNum, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			continue
		}
		versionNum = versionNum + strNum*Pow(1000, int64(len))
		len = len - 1
	}
	return versionNum
}

func Pow(x, n int64) int64 {
	var ret int64 = 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return  "";
}
