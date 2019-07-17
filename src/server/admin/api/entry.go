package api

import (
	"encoding/json"
	log "github.com/astaxie/beego/logs"
	"go_monitor/src/server/admin/handler"
	"io/ioutil"
	"net/http"
)

var (
	apiController = new(ApiController)
)

func InitApi(url string) {
	log.Info("开始初始化api")
	mux := http.NewServeMux()
	handleApi(mux)

	log.Info("监听url:" + url)

	server := &http.Server{
		Addr:    url,
		Handler: mux,
	}

	//go func() {
	err := server.ListenAndServe()
	if err != nil {
		log.Error("监听异常" + err.Error())
	}

	//}()
}

func handleApi(mux *http.ServeMux) {
	mux.HandleFunc("/api/pushMsg", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Info("开始处理推送消息请求")
	//fmt.Println(common.GetLocalIP())
	br, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("报文解析错误")
		WriteJson(w, NewResData(502, "报文解析错误"+err.Error()))
		return
	}
	param := new(handler.Param)
	err = json.Unmarshal(br, &param)
	if err != nil {
		WriteJson(w, NewResData(502, "参数解析错误"+err.Error()))
		return
	}
	err = apiController.PushMsg(param)
	if err != nil {
		WriteJson(w, NewResData(502, "推送消息失败"+err.Error()))
		return
	}
	WriteJson(w, NewResData(200, "成功"))
}

func WriteJson(w http.ResponseWriter, v interface{}) {
	jData, _ := json.Marshal(v)
	w.Write(jData)
}

type ResData struct {
	Code int
	Msg  string
}

func NewResData(code int, msg string) *ResData {
	var resData = new(ResData)
	resData.Code = code
	resData.Msg = msg
	return resData
}
