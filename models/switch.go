package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Switch struct {
	gorm.Model
	Bt     string
	Sbt    string
	Switch string
}

func (s *Switch) Update() {

}
