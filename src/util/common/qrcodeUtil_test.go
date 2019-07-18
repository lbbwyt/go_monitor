package common

import (
	"testing"
)

func TestCreateQrcode(t *testing.T) {
	CreateQrcode("www.baidu.com", "baidu.png")
}

func TestPraseQrcode(t *testing.T) {
	PraseQrcode("baidu.png")
}
