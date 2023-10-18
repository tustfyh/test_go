package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var Pool *redis.Pool

func initPool(address string, maxidle, maxactive int, idtimeout time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxidle,   //最大空闲连接数
		MaxActive:   maxactive, //表示和数据库的连接数
		IdleTimeout: idtimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接代码
			return redis.Dial("tcp", address)
		},
	}
}
