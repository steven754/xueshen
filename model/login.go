package model

import (
	"xueshen/config"
)

type LoginAccount struct {
	BaseEntity
	Account  string `json:"account" gorm:"column:account;unique"`
	Password string `json:"password" gorm:"column:password;"`
}

func CreateLoginAccount(Login *LoginAccount) (err error) {
	err = config.DB.Create(&Login).Error
	return
}

func DeleteLoginAccount(id string) (err error) {
	err = config.DB.Where("id=?", id).Delete(&LoginAccount{}).Error
	return
}
