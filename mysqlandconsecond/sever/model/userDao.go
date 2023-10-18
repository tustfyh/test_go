package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"mysqlandcon/common/message"
)

var (
	MyuserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewRedisConn(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool: pool}
	return userDao
}

// 验证id
func (this *UserDao) getUserByID(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "user", id))
	if err != nil {
		if err == redis.ErrNil { //该值表示哈希表中无该id
			err = ERR_USER_NOTEX
		}
		return nil, err
	}
	//将拿到的res反序列化
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("res json.Unmarshal err=", err)
		return nil, err
	}
	user02 := User{
		UserId:     user.UserId,
		UserName:   user.UserName,
		UserPwd:    user.UserPwd,
		UserStatus: message.UserOnline,
	}
	data, err := json.Marshal(user02)
	_, err = conn.Do("HSet", "user", user02.UserId, string(data))
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 数据库处理登录逻辑
func (this *UserDao) Login(userid int, userpwd string, name string) (user *User, err error) {
	//拿到连接
	conn := this.pool.Get()
	defer conn.Close()
	//验证id
	user, err = this.getUserByID(conn, userid)
	if err != nil {
		return nil, err
	}
	//假如有这个id
	if userpwd != user.UserPwd {
		err = ERR_USER_PWD
		return nil, err
	}
	if name != user.UserName {
		err = ERR_USERNAME
		return nil, err
	}
	if user.UserStatus == message.UserOnline {
		err = ERR_USER_ONLINED
		return nil, err
	}

	return user, nil
}
func (this *UserDao) Signup(userid int, userpwd string, username string) (err error) {
	//拿到连接
	conn := this.pool.Get()
	defer conn.Close()
	//验证id
	_, err = this.getUserByID(conn, userid)
	if err == nil {
		err = ERR_USER_EX
		return err
	}
	//将这个用户写到数据库中
	user := User{
		UserId:     userid,
		UserName:   username,
		UserPwd:    userpwd,
		UserStatus: message.UserOffline,
	}
	data, err := json.Marshal(user)
	_, err = conn.Do("HSet", "user", userid, string(data))
	if err != nil {
		fmt.Println("入库错误")
		return err
	}
	return nil
}
func (this *UserDao) ExitOnline(userid int, pwd string, name string) {
	conn := this.pool.Get()
	defer conn.Close()
	user := User{
		UserId:     userid,
		UserName:   name,
		UserPwd:    pwd,
		UserStatus: 0,
	}
	data, err := json.Marshal(user)
	_, err = conn.Do("HSet", "user", user.UserId, string(data))
	if err != nil {
		return
	}
}
