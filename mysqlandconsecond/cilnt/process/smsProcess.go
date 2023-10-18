package process

import (
	"encoding/json"
	"fmt"
	"mysqlandcon/common/comutils"
	"mysqlandcon/common/message"
)

type SmsProcess struct {
}

func (sp *SmsProcess) sendGroupSes(content string) error {
	//创建一个mes实例
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.UserId = Curuser.UserId
	smsMes.UserName = Curuser.UserName
	smsMes.Content = content
	smsMes.UserStatus = Curuser.UserStatus
	//序列化smsmes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err")
		return err
	}
	mes.Data = string(data)
	//序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err")
		return err
	}
	//发送消息
	tf := &comutils.Transfer{Conn: Curuser.Conn}
	err = tf.PakWrite(data)
	if err != nil {
		fmt.Println("tf.PakWrite(data) err")
		return err
	}
	return nil
}
func (sp *SmsProcess) sendPrivateMes(content string, prvuserid int) error {
	var mes message.Message
	mes.Type = message.SmsPrivateMesType
	var smsprvmes message.SmsPrivateMes
	smsprvmes.Userid = Curuser.UserId
	smsprvmes.Username = Curuser.UserName
	smsprvmes.Content = content
	smsprvmes.UserStatus = Curuser.UserStatus
	smsprvmes.PrivateUserid = prvuserid
	//序列化smsmes
	data, err := json.Marshal(smsprvmes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err")
		return err
	}
	mes.Data = string(data)
	//序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err")
		return err
	}
	tf := &comutils.Transfer{Conn: Curuser.Conn}
	err = tf.PakWrite(data)
	if err != nil {
		fmt.Println("tf.PakWrite(data) err")
		return err
	}
	return nil
}
