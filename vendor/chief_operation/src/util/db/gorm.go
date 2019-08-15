package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Db struct {
	*gorm.DB
	Tx   *gorm.DB
	Data interface{}
}

func (this *Db) Begins() {
	this.Tx = this.DB.Begin()
}

func (this *Db) Commits() error {
	err := this.Tx.Commit().Error
	this.Tx = nil
	return err
}

func (this *Db) Rollbacks() error {
	err := this.Tx.Rollback().Error
	this.Tx = nil
	return err
}

func NewMysql(args string, maxCon int, arr []interface{}) *gorm.DB {
	var con *gorm.DB
	var err error
	con, err = gorm.Open("mysql", args)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, arg = %v the error is '%v'", args, err))
	}
	//设置表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName[:len(defaultTableName)]
	}
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName[:len(defaultTableName)-1]
	//}
	con.LogMode(true) //开启sql日志
	con.DB().SetMaxOpenConns(maxCon)
	idle := maxCon
	if maxCon/3 > 10 {
		idle = maxCon / 3
	}
	con.DB().SetMaxIdleConns(idle)

	//若结构有变，则删除表重新创建
	//dropTable(con, arr...)
	con.AutoMigrate(arr...) //若没有表，自动生成表
	return con
}

func (this *Db) Create() error {
	var err error
	err = this.DB.Create(this.Data).Error
	return err
}

func List(db *gorm.DB, v interface{}, wh ...interface{}) error {
	var err error
	db = db.Model(v)
	if len(wh) == 1 {
		db = db.Where(wh[0])
	} else if len(wh) > 1 {
		db = db.Where(wh[0], wh[1:]...)
	}
	err = db.Scan(v).Error
	return err
}

func Read(db *gorm.DB, v interface{}, wh ...interface{}) error {
	var err error
	if len(wh) == 1 {
		db = db.Where(wh[0])
	} else if len(wh) > 1 {
		db = db.Where(wh[0], wh[1:]...)
	}
	err = db.First(v).Error
	return err
}

func Del(db *gorm.DB, v interface{}, wh ...interface{}) error {
	var err error
	if len(wh) == 1 {
		db = db.Where(wh[0])
	} else if len(wh) > 1 {
		db = db.Where(wh[0], wh[1:]...)
	}
	err = db.Delete(v).Error
	return err
}

func Count(db *gorm.DB, v interface{}, wh ...interface{}) (int64, error) {
	var err error
	db = db.Model(v)
	if len(wh) == 1 {
		db = db.Where(wh[0])
	} else if len(wh) > 1 {
		db = db.Where(wh[0], wh[1:]...)
	}
	var count int64 = 0
	err = db.Count(&count).Error
	return count, err
}

func Update(db *gorm.DB, f interface{}, key string, value interface{}, wh ...interface{}) error {
	var err error
	db = db.Model(f)
	if len(wh) == 1 {
		db = db.Where(wh[0])
	} else if len(wh) > 1 {
		db = db.Where(wh[0], wh[1:]...)
	}
	err = db.Update(key, value).Error
	return err
}

//更新多个值 db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
func Updates(db *gorm.DB, f interface{}, value interface{}, wh ...interface{}) error {
	var err error
	db = db.Model(f)
	if len(wh) == 1 {
		db = db.Where(wh[0])
	} else if len(wh) > 1 {
		db = db.Where(wh[0], wh[1:]...)
	}
	err = db.Updates(value).Error
	return err
}

func Save(db *gorm.DB, v interface{}) error {
	var err error
	err = db.Save(v).Error
	return err
}
