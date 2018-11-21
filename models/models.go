package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/oumeniOS/go-gin-blog/pkg/setting"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	gorm.Model
}

var db *gorm.DB

func Setup() {
	var err error

	//打开数据库
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("Fail gorm.Open %v", err)
	}

	//设置数据库名称前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	//默认单表
	db.SingularTable(true)
	//设置最大空闲链接
	db.DB().SetMaxIdleConns(10)
	//设置最大链接数
	db.DB().SetMaxOpenConns(100)
}

//关闭数据库
func CloseDB() {
	defer db.Close()
}
