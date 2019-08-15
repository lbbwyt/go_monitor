package common

type Param struct {
	Org string  `json:"Org"` //单位名称 数字高新
	AppName string `json:"AppName"` //应用名称 monitor
	ModuleName string  `json:"ModuleName"`// 模块名称 智慧消防
	Code  string `json:"Code"` // 错误码
	Msg   string  `json:"Msg"`// 系统错误消息，较详细
	Error string  `json:"Error"`//自定义错误消息，
}

