package dao

import (
	"chief_operation/src/server/admin/config"
	"chief_operation/src/util/db"
	"github.com/jinzhu/gorm"
	log "github.com/astaxie/beego/logs"
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