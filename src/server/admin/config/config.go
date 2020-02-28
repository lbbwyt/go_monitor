package config

import (
	"encoding/json"
	"fmt"
	log "github.com/astaxie/beego/logs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go_monitor/src/util/form"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	Conf = new(form.CommonConf)
)

func InitConfig(test string) {
	log.Info("开始加载配置")
	ReadConf(Conf, test)
	//从配置文件中加载数据库配置,配置文件变更后，重启服务。
	log.Info("加载配置完成")
}

func InitVipConfig() {
	viper.SetConfigName("conf")
	viper.AddConfigPath(GetCurrentDirectory() + "/config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("配置文件不存在")
		} else {
			log.Error("读取配置文件失败")
		}
		return
	}

	var confWxMap = viper.GetStringMapString("weixin")
	var confDDMap = viper.GetStringMapString("dingding")
	var confMysqlMap = viper.GetStringMapString("mysql")
	var confSmsMap = viper.GetStringMapString("sms")
	var (
		path   string = confDDMap["path"]
		send   int
		limit  int
		people []string = make([]string, 0)
	)
	agentid, _ := strconv.Atoi(confWxMap["agentid"])
	var weixin = form.NewWeixin(agentid, confWxMap["corpid"], confWxMap["corpsecret"])
	var maxCon, _ = strconv.Atoi(confMysqlMap["max_con"])
	var mysql = form.NewMysql(confMysqlMap["host"], confMysqlMap["db"], maxCon)

	send, _ = strconv.Atoi(confDDMap["send"])
	limit, _ = strconv.Atoi(confDDMap["limit_send"])
	peopleStr := confDDMap["people"]
	people = append(people, peopleStr)
	log.Info("peopleStr" + peopleStr)
	var dingding = form.NewDingDing(path, send, people, limit)

	smsEnabled, _ := strconv.Atoi(confSmsMap["enabled"])

	Sms := form.NewSms(smsEnabled, confSmsMap["host"], confSmsMap["phone"])

	Conf = form.NewCommonConf(*mysql, *weixin, *dingding, *Sms)

	log.Info("加载配置完成")
	go WatchConfig()

}

func WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:")
		InitVipConfig()
	})
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
