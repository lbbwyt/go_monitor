package api

import (
	"fmt"
	log "github.com/astaxie/beego/logs"
	"go_monitor/src/server/admin/config"
	"go_monitor/src/server/admin/dao"
	"go_monitor/src/server/admin/entity"
	"go_monitor/src/server/admin/handler"
	emap "go_monitor/src/util/ExpiredMap"
	"go_monitor/src/util/common"
	"time"
)

type ApiController struct{}

var (
	cacheMap = emap.NewExpiredMap()
)

func NewApiController() *ApiController {
	d := new(ApiController)
	return d
}

func (this *ApiController) PushMsg(param *handler.Param) error {
	log.Info("推送消息：" + common.Obj2JsonStr(param))
	//消息推送到微信
	var content = fmt.Sprintf("异常单位：%v, 错误提示为：%v, 具体信息为：%v",
		param.Org, param.Error, param.Msg)
	log.Info("************" + content)
	found, value := cacheMap.Get(content)
	if found {
		log.Info(value.(string) + " already alarm, ttl：")
		log.Info(cacheMap.TTL(content))
		return nil
	}
	cacheMap.Set(content, content, int64(60*60))
	log.Info("start alarm")
	//消息推送到钉钉
	go handler.Add(param)

	go handler.PushWxMsg(content)

	//推送短信消息
	if config.Conf.Sms.Enabled == 1 {
		go handler.PushSmsMsg(content, config.Conf.Sms.Phone)
	}

	var alarmMsg = new(entity.AlarmMsg)
	alarmMsg.AppName = param.AppName
	alarmMsg.CreateTime = time.Now()
	alarmMsg.StartTime = time.Now()
	alarmMsg.ErrMsg = param.Msg
	alarmMsg.Org = param.Org
	alarmMsg.Title = param.Error
	alarmMsg.ModuleName = param.ModuleName
	//写数据到数据库
	err := dao.MysqlCon.Save(alarmMsg).Error
	if err != nil {
		log.Error("写入消息失败" + err.Error())
		return err
	}
	return nil
}
