package ExpiredMap

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"testing"
	"time"
)

var (
	cacheMap = NewExpiredMap()
)

func TestExpiredMap_Get(t *testing.T) {

	var content = fmt.Sprintf("异常单位：%v, 错误提示为：%v, 具体信息为：%v",
		"qweqwe", "qweqwe", "qweqwe")

	found, _ := cacheMap.Get(content)
	if found {
		logs.Info("already alarm")
		return
	}
	cacheMap.Set(content, 1, int64(60*60))
	logs.Info("start alarm")

	time.Sleep(time.Second * 6)
	fmt.Println(cacheMap.Get(content))
	fmt.Println(cacheMap.TTL(content))
}
