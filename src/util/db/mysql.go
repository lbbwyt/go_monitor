package db

import (
	log "github.com/astaxie/beego/logs"
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	MysqlCon    *gorm.DB
	MysqlLogCon *gorm.DB
	MysqlActCon *gorm.DB
)

func InitMysql(host, db string, maxCon int, arr ...interface{}) {
	log.Info("init mysql")
	//arg := host + "/" + db + "?charset=utf8&parseTime=True&loc=Local"
	arg := fmt.Sprintf("%v/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", host, db)
	MysqlCon = NewMysql(arg, maxCon, arr)
	log.Info("初始化mysql完成")
}

func InitMysqlLog(host, db string, maxCon int, arr ...interface{}) {
	log.Info("init mysql")
	//arg := host + "/" + db + "?charset=utf8&parseTime=True&loc=Local"
	arg := fmt.Sprintf("%v/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", host, db)
	MysqlLogCon = NewMysql(arg, maxCon, arr)
	log.Info("初始化mysql log完成")
}

func InitMysqlAct(host, db string, maxCon int, arr ...interface{}) {
	log.Info("init mysql")
	//arg := host + "/" + db + "?charset=utf8&parseTime=True&loc=Local"
	arg := fmt.Sprintf("%v/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", host, db)
	MysqlActCon = NewMysql(arg, maxCon, arr)
	log.Info("初始化mysql完成")
}
