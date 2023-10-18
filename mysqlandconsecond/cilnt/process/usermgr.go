package process

import (
	"encoding/json"
	"fmt"
	"mysqlandcon/cilnt/modelcli"
	"mysqlandcon/common/message"
	"mysqlandcon/sever/model"
)

var OnlineUsers map[int]*model.User = make(map[int]*model.User, 10)
var Curuser modelcli.CurUser

func Output02() {
	for id, user := range OnlineUsers {
		if user.UserStatus == message.UserOnline {
			fmt.Printf("用户%d在线\n", id)
		} else if user.UserStatus == message.UserOffline {
			//处理逻辑
			fmt.Printf("用户%d下线\n", id)
		}
	}
}
func Output() {
	for id, _ := range OnlineUsers {
		fmt.Printf("用户%d在线\n", id)
	}
}

// 编写一个方法，处理notify
func UpdateUserStatus(notify message.NotifyUserStatusMes) {
	user, ok := OnlineUsers[notify.Userid]
	if ok == false {
		user = &model.User{
			UserId: notify.Userid,
		}
	}
	user.UserStatus = notify.Status
	OnlineUsers[notify.Userid] = user
	Output02()
}

// 处理下线
func OutputExitMes(mes *message.Message) {
	var exitresmes message.DownOnlineMes
	err := json.Unmarshal([]byte(mes.Data), &exitresmes)
	if err != nil {
		return
	}
	for i, _ := range OnlineUsers {
		if i == exitresmes.Userid {
			delete(OnlineUsers, i)
		}
	}
}
