package config

import (
	"chief_operation/src/util/form"
	"encoding/json"
	"fmt"
	log "github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)



type Config struct {
	*form.PortConf
}

var (
	Conf =  new(Config)
)

func InitConfig(test string) {
	log.Info("开始加载配置")
	ReadConf(Conf, test)

	//从配置文件中加载数据库配置,配置文件变更后，重启服务。
	log.Info("加载配置完成")
}



func ReadConf(v interface{}, test string) {
	fmt.Println("当前路径为：" + GetCurrentDirectory())
	var pix = GetCurrentDirectory()
	if test != "" {
		pix = test
	}
	data, err := ioutil.ReadFile(pix + "/config/conf.json")
	if err != nil {
		panic(err)
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Error("初始化配置异常")
		return
	}
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("无法获取当前文件信息")
	}
	return strings.Replace(dir, "\\", "/", -1)
}