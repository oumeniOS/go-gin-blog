package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/oumeniOS/go-gin-blog/pkg/setting"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	gorm.Model
}

var db *gorm.DB

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database' %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	//打开数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))
	if err != nil {
		log.Fatalf("Fail gorm.Open %v", err)
	}

	//设置数据库名称前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//默认单表
	db.SingularTable(true)
	//设置最大空闲链接
	db.DB().SetMaxIdleConns(10)
	//设置最大链接数
	db.DB().SetMaxOpenConns(100)

	//db.Callback().Create().Replace("gorm:after_create", UpdateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", UpdateTimeStampForUpdateCallback)
}

//关闭数据库
func CloseDB() {
	defer db.Close()
}

//func UpdateTimeStampForCreateCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		nowTime := time.Now().Unix()
//		//set CreateOn
//		if createTimeField, ok := scope.FieldByName("CreateOn"); ok == true {
//			if createTimeField.IsBlank {
//				createTimeField.Set(nowTime)
//			}
//		}
//
//		//set ModifiedOn
//		if modifiedTimeField, ok := scope.FieldByName("ModifiedOn"); ok == true {
//			if modifiedTimeField.IsBlank {
//				modifiedTimeField.Set(nowTime)
//			}
//		}
//	}
//}
//
//func UpdateTimeStampForUpdateCallback(scope *gorm.Scope) {
//	nowTime := time.Now().Unix()
//	if _, ok := scope.Get("gorm:update_column"); !ok {
//		scope.SetColumn("ModifiedOn", nowTime)
//	}
//}






































