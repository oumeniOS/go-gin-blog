package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/EDDYCJY/gin-blog/pkg/setting"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
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
}



//关闭数据库
func CloseDB() {
	defer db.Close()
}
