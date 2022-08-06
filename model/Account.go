package model

import "xueshen/config"

type SignAccount struct {
	BaseEntity
	Account  string `json:"account" gorm:"column:account;unique" binding:"required"`
	Password string `json:"password" gorm:"column:password" binding:"required"`
}

func CreateSignAccount(Sign *SignAccount) (err error) {
	err = config.DB.Create(&Sign).Error
	return
}

//func CheckAccount(Sign *SignAccount) (err error) {
//	account := config.DB.First(Sign.Account)
//	password := config.DB.Where("password=?)
//	if account ==
//	err = config.DB.Create(&Sign).Error
//	return
//}
