package process

import (
	"encoding/json"
	"fmt"
	"mysqlandcon/common/comutils"
	"mysqlandcon/common/message"
	"mysqlandcon/sever/model"
	"net"
)

type UserProcess struct {
	//字段
}

func (up *UserProcess) Login(userid int, pwd string, name string) (err error) {
	//1.连接到服务器端
	conn, err := net.Dial("tcp", "192.168.31.9:9000")
	if err != nil {
		fmt.Printf("net.Dial err=%v\n", err)
		return err
	}
	tf := &comutils.Transfer{Conn: conn}
	//defer conn.Close()//如果关闭会有意想不到的错误
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LogInMesType
	//3.进行序列化
	var loginmes message.LogInMes
	loginmes.UserPwd = pwd
	loginmes.UserId = userid
	loginmes.Username = name
	data, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return err
	}
	mes.Data = string(data)
	//对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return err
	}
	//发送消息
	err = tf.PakWrite(data)
	if err != nil {
		return err
	}
	//读取消息-----读取返回的登录消息
	ResMes, err := tf.PakRed()
	if err != nil {
		return err
	}
	//解包
	var LoginResMes message.LoginResMes
	err = json.Unmarshal([]byte(ResMes.Data), &LoginResMes)
	if err != nil {
		return err
	}

	if LoginResMes.Code == 200 {
		Curuser.Conn = conn
		Curuser.UserId = userid
		Curuser.UserName = name
		Curuser.UserStatus = message.UserOnline
		for _, v := range LoginResMes.UsersId {
			if v == userid {
				continue
			} else {
				fmt.Printf("\t用户id=%d的在线\n", v)
			}
			user := &model.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			OnlineUsers[user.UserId] = user
		}
		print("\n")
		//这里需要起一个协程来保持与客户端的通讯
		go keepSeverMes(conn)
		for {
			ShowMenu(userid, tf, name, pwd)
		}
	} else {
		fmt.Println(LoginResMes.Err)
	}
	return nil
}

func (this *UserProcess) Signup(userid int, pwd string, name string) (err error) {
	//1.连接到服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Printf("net.Dial err=%v\n", err)
		return err
	}
	tf := &comutils.Transfer{Conn: conn}
	//defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.SignupMesType
	//3.进行序列化
	var signupmes message.SignupMes
	signupmes.UserPwd = pwd
	signupmes.UserId = userid
	signupmes.UserName = name
	data, err := json.Marshal(signupmes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return err
	}
	mes.Data = string(data)
	//对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return err
	}
	err = tf.PakWrite(data)
	if err != nil {
		return err
	}
	//解包
	ResMes, err := tf.PakRed()
	if err != nil {
		return err
	}
	var signupResMes message.SignupResMes
	err = json.Unmarshal([]byte(ResMes.Data), &signupResMes)
	if signupResMes.Code == 200 {
		fmt.Println("注册成功，请进行登录")
	} else {
		fmt.Println(signupResMes.Err)
	}
	return nil
}
