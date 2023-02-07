package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func ConnRedis() {

	// 连接超时时间
	optionTimeout := redis.DialConnectTimeout(time.Second * 10)

	// 连接哪个数据库 默认是第0个数据库
	optionDb := redis.DialDatabase(1)

	// 连接的密码（如果设置了的话）
	optionPwd := redis.DialPassword("")

	// 2.连接redis服务 第一个参数是network类型 redis是tcp/ip协议，第二个参数是redisHost:redisPort 第三个参数为连接配置
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", optionTimeout, optionDb, optionPwd)
	if err != nil {
		fmt.Println("conn redis server err,", err.Error())
		return
	}

	// 3.处理完之后关闭redis连接
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//// 4.redis set 命令操作 用Do方法发送redis操作命令，返回操作的结果
	//replyRes, err := conn.Do("set", "username", "点滴课堂")
	//
	//// 使用redis包中的方法解析返回结果
	//str, _ := redis.String(replyRes, err)
	//fmt.Println(str) //OK
	//
	//// 5.redis get 命令操作
	//replyRes, err = conn.Do("get", "username")
	//
	//// 使用redis包中的方法解析返回结果
	//str, _ = redis.String(replyRes, err)
	//fmt.Println(str) //点滴课堂
	//
	//// 6.redis list 命令操作
	//replyRes, err = conn.Do("lpush", "msg-list", "msg1")
	//str, _ = redis.String(replyRes, err)
	//fmt.Println(str) //空字符串
	//
	//replyRes, err = conn.Do("lpop", "msg-list")
	//str, _ = redis.String(replyRes, err)
	//fmt.Println(str) //msg1
}
