package models

import (
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"gorm.io/datatypes"
)

type Content struct {
	Id      uint           `gorm:"primary_key" json:"vid"`
	Mid     uint           `json:"id"`
	Bt      string         `gorm:"type:varchar(40);not null" json:"bt"`
	Sbt     string         `gorm:"type:varchar(40);not null" json:"sbt"`
	Source  string         `gorm:"type:varchar(4);not null" json:"source"`
	Service string         `gorm:"type:varchar(8);not null" json:"service"`
	Op      string         `gorm:"type:varchar(8);not null" json:"op"`
	Status  string         `gorm:"type:varchar(8);not null" json:"status"`
	Reason  string         `gorm:"type:varchar(80)" json:"reason"`
	Data    datatypes.JSON `gorm:"type:json" json:"data"`
}

//自定义表名
func (Content) TableName() string {
	return "oplog"
}

func GetOplogs(start, limit int) []Content {
	oplogs := []Content{}
	DB.Where("id >= ? AND id < ?", start, start+limit).Find(&oplogs)
	return oplogs
}

func GetDevId() uint {
	var c Content
	r := DB.Last(&c)
	if r.Error != nil {
		return 0
	}
	return c.Id
}

func GetMsId() uint {
	var c Content
	r := DB.Order("mid desc, mid").First(&c)
	if r.Error != nil {
		return 0
	}
	return c.Mid
}

func SetMsId(mid uint) {
	var c Content
	r := DB.Order("mid desc, mid").First(&c)
	if r.Error != nil {
		return
	}
	DB.Model(&c).Update("mid", mid)
}
