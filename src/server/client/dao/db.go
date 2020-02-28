package dao

import (
	log "github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"go_monitor/src/server/client/config"

	"go_monitor/src/util/db"
)

var (
	MysqlCon *gorm.DB
)

func InitMysql() {
	log.Info("init mysql")
	host := config.Conf.Mysql.Host
	dbName := config.Conf.Mysql.Db
	arg := host + "/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	var maxCon = 20
	var arr = []interface{}{}
	MysqlCon = db.NewMysql(arg, maxCon, arr)
	log.Info("初始化MysqlCon完成")
}
