package model

import (
	"errors"
	"xueshen/config"
)

type SignAccount struct {
	BaseEntity
	Account  string `json:"account" gorm:"column:account" binding:"required"`
	Password string `json:"password" gorm:"column:password" binding:"required"`
}

func CreateSignAccount(Sign *SignAccount) (err error) {
	err = CheckAccountExist(Sign)
	if err != nil {
		return err
	}
	err = config.DB.Create(&Sign).Error
	if err != nil {
		return err
	}
	return
}

func CheckAccountExist(Sign *SignAccount) (err error) {
	result := config.DB.Table("sign_accounts").Where("account=?", Sign.Account).First(&Sign)
	if result.RowsAffected == 0 {
		return
	} else {
		return errors.New("该账号已注册")
	}
}
