package main

import (
	log "github.com/astaxie/beego/logs"
	dispatcher "go_monitor/src/server/short_url/dispatcher"
	"net/http"
)

/*
 * 短地址服务，
 */
func main() {
	log.Info("Enter main")
	Init()
}

//初始化
func Init() {
	log.Info(">>>>开始初始化")
	http.HandleFunc("/", dispatcher.SayHello) //注册URI路径与相应的处理函数
	log.Info("【默认项目】服务启动成功 监听端口 80")
	er := http.ListenAndServe("0.0.0.0:80", nil)
	if er != nil {
		log.Error("ListenAndServe: ", er)
	}
}
