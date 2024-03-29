package model

import (
	"errors"
	"fmt"
	"time"
	"xueshen/config"
)

type SignAccount struct {
	BaseEntity
	Account  string `json:"account" gorm:"column:account" binding:"required"`
	Password string `json:"password" gorm:"column:password"`
}

type AccountList struct {
	Account  string `json:"account"`
	PageNo   uint   `json:"PageNo" gorm:"default:1"`
	PageSize uint   `json:"PageSize" gorm:"default:10"`
	Total    uint   `json:"Total"`
}

type QueryAccount struct {
	Account string `json:"account" gorm:"column:account"`
}

type AccountListRsp struct {
	AccountList
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
	result := config.DB.Model(&Sign).Where("account=?", Sign.Account).First(&Sign)
	if result.RowsAffected == 0 {
		return
	} else {
		return errors.New("该账号已注册")
	}
}

// 判断密码是否正确
func CheckPasswordCorrect(Sign *SignAccount) (err error) {
	result := config.DB.Model(&Sign).Where("account=? AND password=?", Sign.Account, Sign.Password).Find(&Sign)
	fmt.Println(result)
	if result.RowsAffected == 0 {
		return errors.New("密码错误")
	} else {
		return
	}
}

// 修改密码
func UpdateAccountPassword(Sign *SignAccount) (err error) {
	//err = CheckAccountExist(Sign)
	//if err != nil {
	//	return err
	//}
	Sign.UpdateTime = time.Now()
	//config.DB.Model(&Sign).Where("account=?", Sign.Account).Update("password", Sign.Password)
	err = config.DB.Model(&Sign).Where("account=?", Sign.Account).Update("password", Sign.Password).Error
	if err != nil {
		return err
	}
	return
}

func QueryAccountList(queryAccount *SignAccount) (rsp *[]string, err error) {
	result := make([]string, 10)
	err = config.DB.Model(&queryAccount).Where("account LIKE ?", "%"+queryAccount.Account+"%").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
