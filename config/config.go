package config

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// 切换环境
var terminal = "mac" //win，mac

// windows电脑
var (
	WinHost            = "192.168.199.128:3306"
	WinUser            = "root"
	WinPassWord string = "123456"
	WinDataBase        = "xueshen"
	WinUrl             = "localhost:9000"
)

// mac电脑
var (
	MacHost     = "localhost:3306"
	MacUser     = "root"
	MacPassWord = "123456"
	MacDataBase = "xueshen"
	MacUrl      = "localhost:9000"
)

func init() {
	Getconf()
}

func Getconf() (c map[string]interface{}) {
	conf := make(map[string]interface{})
	if terminal == "win" {
		conf["Host"] = WinHost
		conf["User"] = WinUser
		conf["PassWord"] = WinPassWord
		conf["DataBase"] = WinDataBase
		conf["Url"] = WinUrl
	} else if terminal == "mac" {
		conf["Host"] = MacHost
		conf["User"] = MacUser
		conf["PassWord"] = MacPassWord
		conf["DataBase"] = MacDataBase
		conf["Url"] = MacUrl
	} else {
		fmt.Println("请输入正确的termianl类型")
		panic("termianl类型只能是小写的win和mac")
	}
	return conf
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
