package models

import (
//"gorm.io/driver/mysql"
//"gorm.io/gorm"
)

type Oplog struct {
	Id      uint   `gorm:"primary_key" json:"vid"`
	Mid     uint   `json:"id"`
	Bt      string `gorm:type:varchar(40);not null" json:"bt"`
	Sbt     string `gorm:type:varchar(40);not null" json:"sbt"`
	Source  string `gorm:type:varchar(4);not null" json:"source"`
	Service string `gorm:type:varchar(8);not null" json:"service"`
	Op      string `gorm:type:varchar(8);not null" json:"action"`
	Reason  string `json:"reason"`
	Data    map[string]interface{}
}
