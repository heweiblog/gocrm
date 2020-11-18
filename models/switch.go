package models

import (
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Switch struct {
	gorm.Model `json:"-"`
	Bt         string `gorm:"type:varchar(40);not null" json:"bt"`
	Sbt        string `gorm:"type:varchar(40);not null" json:"sbt"`
	Service    string `gorm:"type:varchar(8);not null" json:"service"`
	Switch     string `gorm:"type:varchar(8);not null" json:"switch"`
}

func () GetSwitchs() []Switch {
	switchs := []Switch{}
	db.Find(&switchs)
	return switchs
}

func (s *Switch) UpdateSwitch(bt, sbt, service, status, string) {
}
