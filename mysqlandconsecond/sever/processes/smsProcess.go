package processes

import (
	"encoding/json"
	"mysqlandcon/common/comutils"
	"mysqlandcon/common/message"
)

type SmsProcess struct {
}

func (sp *SmsProcess) TurnGroupMes(mes *message.Message) {
	var smsmes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsmes)
	if err != nil {
		return
	}
	data, err := json.Marshal(mes)
	for id, up := range Usermgr.UserOnline {
		if id == smsmes.UserId {
			continue
		}
		sendMes(up, data)
	}
}
func (sp *SmsProcess) TurnPrivateMes(mes *message.Message) {
	var smsPrivateMes message.SmsPrivateMes
	err := json.Unmarshal([]byte(mes.Data), &smsPrivateMes)
	if err != nil {
		return
	}
	data, err := json.Marshal(mes)
	for id, up := range Usermgr.UserOnline {
		if id == smsPrivateMes.Userid {
			continue
		}
		if id == smsPrivateMes.PrivateUserid {
			sendMes(up, data)
		}
	}

}
func sendMes(up *UserProcess, data []byte) {
	tf := &comutils.Transfer{Conn: up.Conn}
	//发送原始数据
	err := tf.PakWrite(data)
	if err != nil {
		return
	}
}
