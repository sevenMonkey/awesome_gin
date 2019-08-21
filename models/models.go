package models

import (
	"awesome_gin/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Model struct {
	ID      int `gorm:"primary_key" json:"id"`
	Created int `json:"created" gorm:"column:created"`
	Updated int `json:"updated" gorm:"column:updated"`
}

var db *gorm.DB

func Setup()  {
	var err error
	db, err = gorm.Open(setting.DBSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DBSetting.User,
			setting.DBSetting.Password,
			setting.DBSetting.Host,
			setting.DBSetting.DbName))
	if err != nil{
		log.Fatalf("models.setup err:%v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("Created"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("Updated"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("Updated", time.Now().Unix())
	}
}

func CloseDB()  {
	if err :=  db.Close(); err != nil{
		log.Fatalf("models.close err:%v", err)
	}
}
