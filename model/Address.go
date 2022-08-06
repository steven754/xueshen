package model

import "xueshen/config"

type Address struct {
	BaseEntity
	Account  string `json:"account" gorm:"column:account;unique"`
	Password string `json:"password" gorm:"column:password;"`
}

func AddressInfo(Addr *Address) (err error) {
	err = config.DB.Create(&Addr).Error
	return
}
