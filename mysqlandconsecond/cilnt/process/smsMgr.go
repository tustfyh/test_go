package process

import (
	"encoding/json"
	"fmt"
	"mysqlandcon/common/message"
)

func OutputGroupMes(mes *message.Message) {
	var smsmes message.SmsMes

	err := json.Unmarshal([]byte(mes.Data), &smsmes)
	if err != nil {
		return
	}
	fmt.Printf("用户id为%d,姓名为%s进行世界广播:%s\n", smsmes.UserId, smsmes.UserName, smsmes.Content)
}
func OutputPrivateMes(mes *message.Message) {

	var smsPrivateMes message.SmsPrivateMes

	err := json.Unmarshal([]byte(mes.Data), &smsPrivateMes)
	if err != nil {
		return
	}
	fmt.Printf("用户id为%d,姓名为%s对你偷偷私发:%s\n", smsPrivateMes.Userid, smsPrivateMes.Username, smsPrivateMes.Content)
}
