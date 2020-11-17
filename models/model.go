package models

import (
	"fmt"
	"gocrm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

var DB *gorm.DB

func init() {
	sql := config.Conf.Sql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", sql.User, sql.Pass, sql.Ip, strconv.Itoa(sql.Port), sql.Database)
	//因为DB为全局变量此处必须显式声明error
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(&Product{})
}
