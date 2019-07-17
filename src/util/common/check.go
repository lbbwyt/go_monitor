package common

import (
	"github.com/astaxie/beego/logs"
)

func CheckGoPanic() {
	if r := recover(); r != nil {
		logs.Error("从异常中恢复")
	}
}
