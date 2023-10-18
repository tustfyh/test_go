package main

import (
	"encoding/json"
	"fmt"
	"mysqlandcon/common/comutils"
	"mysqlandcon/common/message"
	"mysqlandcon/sever/processes"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) severProcessMes(mes *message.Message) (err error) {
	//看看服务器收到了的信息
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LogInMesType:
		{
			//处理登录逻辑
			up := &processes.UserProcess{Conn: this.Conn}
			err = up.SeverProcessLogin(mes)
		}
	case message.SignupMesType:
		{
			//处理注册的逻辑
			up := &processes.UserProcess{Conn: this.Conn}
			err = up.SeverProcessSignup(mes)
		}
	case message.SmsMesType:
		{
			//处理消息的逻辑
			sp := &processes.SmsProcess{}
			sp.TurnGroupMes(mes)
		}
	case message.SmsPrivateMesType:
		{
			//处理私法消息的逻辑
			sp := &processes.SmsProcess{}
			sp.TurnPrivateMes(mes)
		}
	case message.DownOnlineMesType:
		{
			//处理下线的逻辑
			var exitonline message.DownOnlineMes
			err = json.Unmarshal([]byte(mes.Data), &exitonline)
			if err != nil {
				return
			}
			Eo := &processes.ExitOnline{}
			up := &processes.UserProcess{Conn: this.Conn}
			processes.Usermgr.UpdateOnlineUser(up, exitonline.UserStatus)
			up.NotifyOthers(exitonline.Userid, exitonline.UserStatus)
			Eo.ProcessExitOnline(mes)
		}
	default:
		{
			fmt.Println("无法处理该类型消息")
		}

	}
	return err
}
func (this *Processor) process02() error {
	//接受客户端发送的消息
	for {
		tf := &comutils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.PakRed()
		if err != nil {
			return err
		}
		err = this.severProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
