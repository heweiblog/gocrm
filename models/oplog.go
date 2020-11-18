package models

import (
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"gorm.io/datatypes"
)

type Oplog struct {
	Id      uint           `gorm:"primary_key" json:"vid"`
	Mid     uint           `json:"id"`
	Bt      string         `gorm:"type:varchar(40);not null" json:"bt"`
	Sbt     string         `gorm:"type:varchar(40);not null" json:"sbt"`
	Source  string         `gorm:"type:varchar(4);not null" json:"source"`
	Service string         `gorm:"type:varchar(8);not null" json:"service"`
	Op      string         `gorm:"type:varchar(8);not null" json:"action"`
	Reason  string         `gorm:"type:varchar(80);not null" json:"reason"`
	Data    datatypes.JSON `gorm:"type:json" json:"data"`
}

func GetOplogs(start, limit int) []Oplog {
	oplogs := []Oplog{}
	DB.Where("id >= ? AND id < ?", start, start+limit).Find(&oplogs)
	return oplogs
}
