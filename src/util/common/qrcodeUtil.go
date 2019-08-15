package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/skip2/go-qrcode"
	t_qrcode "github.com/tuotoo/qrcode"
	"os"
)

func CreateQrcode(content string, filename string) {
	err := qrcode.WriteFile(content, qrcode.Medium, 256, filename)
	if err != nil {
		logs.Info("生产二维码失败" + err.Error())
	}
}

func PraseQrcode(path string) {
	fi, err := os.Open(path)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := t_qrcode.Decode(fi)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	logs.Info(qrmatrix.Content)
}
