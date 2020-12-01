package models

import (
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Register struct {
	gorm.Model `json:"-"`
	ConfUrl    string `gorm:"type:varchar(80);not null"`
	TaskUrl    string `gorm:"type:varchar(80);not null"`
	Module     string `gorm:"type:varchar(40);not null"`
}

func (Register) TableName() string {
	return "register"
}

func GetRegisters() []Register {
	registers := []Register{}
	DB.Find(&registers)
	return registers
}

//func (s *Register) UpdateRegister(bt, sbt, service, status, string) {
//}
