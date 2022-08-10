package model

import (
	"errors"
	"time"
	"xueshen/config"
)

type SignAccount struct {
	BaseEntity
	Account  string `json:"account" gorm:"column:account" binding:"required"`
	Password string `json:"password" gorm:"column:password" binding:"required"`
}

// 创建账号
func CreateSignAccount(Sign *SignAccount) (err error) {
	err = CheckAccountExist(Sign)
	if err != nil {
		return err
	}
	Sign.CreateTime = time.Now()
	Sign.UpdateTime = time.Now()
	err = config.DB.Create(&Sign).Error
	if err != nil {
		return err
	}
	return
}

// 检查账号是否存在
func CheckAccountExist(Sign *SignAccount) (err error) {
	result := config.DB.Table("sign_accounts").Where("account=?", Sign.Account).First(&Sign)
	if result.RowsAffected == 0 {
		return
	} else {
		return errors.New("该账号已注册")
	}
}

// 判断密码是否正确
func CheckPasswordCorrect(Sign *SignAccount) (err error) {
	result := config.DB.Table("sign_accounts").Where("account=? AND password=?", Sign.Account, Sign.Password)
	if result.RowsAffected == 0 {
		return errors.New("密码错误")
	} else {
		return
	}
}

// 修改密码
func UpdateAccountPassword(Sign *SignAccount) (err error) {
	err = CheckAccountExist(Sign)
	if err != nil {
		return err
	}
	Sign.UpdateTime = time.Now()
	err = config.DB.Update("password=?", Sign.Password).Where("account=?", Sign.Account).Error
	if err != nil {
		return err
	}
	return
}
