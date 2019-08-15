package config

import (
	"encoding/json"
	"fmt"
	log "github.com/astaxie/beego/logs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Org     string   `json:"org"`
	Url     string   `json:"url"`
	Logpath string   `json:"logpath"`
	Ports   []string `json:"ports"`
}

var (
	Conf = new(Config)
)

func InitConfig(test string) {
	log.Info("开始加载配置")
	ReadConf(Conf, test)
	//从配置文件中加载数据库配置,配置文件变更后，重启服务。
	log.Info("加载配置完成")
}

func InitVipConfig() {
	viper.SetConfigName("conf")                            // name of config file (without extension)
	viper.AddConfigPath(GetCurrentDirectory() + "/config") // path to look for the config file in
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("配置文件不存在")
		} else {
			log.Error("读取配置文件失败")
		}
		return
	}
	var url = viper.GetString("url")
	Conf.Url = url
	var logpath = viper.GetString("logpath")
	Conf.Logpath = logpath

	Conf.Org = viper.GetString("org")
	var prots []string
	prots = viper.GetStringSlice("ports")
	Conf.Ports = prots
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
