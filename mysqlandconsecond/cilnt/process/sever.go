package process

import (
	"encoding/json"
	"fmt"
	"io"
	"mysqlandcon/common/comutils"
	"mysqlandcon/common/message"
	"net"
	"os"
)

func ShowMenu(userid int, tf *comutils.Transfer, name string, pwd string) {
	var key int
	fmt.Println("-----------------恭喜你登录成功------------------------")
	fmt.Println("\t\t\t1.显示在线用户")
	fmt.Println("\t\t\t2.发送消息")
	fmt.Println("\t\t\t3.退出")
	fmt.Println("\t\t\t请选择<1~3>")
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		{
			fmt.Println("显示在线用户")
			Output()
		}
	case 2:
		{
			fmt.Println("发送消息")
			showSmsTypeMenu()

		}
	case 3:
		{
			ExitOnline(userid, tf, name, pwd)
			fmt.Println("正在退出系统")
			os.Exit(0)
		}
	default:
		{
			fmt.Println("输入错误,请重新输入")
		}
	}
}
func keepSeverMes(conn net.Conn) {
	tf := &comutils.Transfer{Conn: conn}
	for {
		mes, err := tf.PakRed()
		if err != nil {
			if err == io.EOF {
			}
			fmt.Println(err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType: //通知别人上下线的消息
			{
				var notify message.NotifyUserStatusMes
				err = json.Unmarshal([]byte(mes.Data), &notify)
				if err != nil {
					return
				}
				UpdateUserStatus(notify)
			}
		case message.SmsMesType:
			{
				OutputGroupMes(&mes)
			}
		case message.SmsPrivateMesType:
			{
				OutputPrivateMes(&mes)
			}
		case message.DownOnlineMesType:
			{
				//处理逻辑
				OutputExitMes(&mes)
			}

		default:
			fmt.Println("类型无法匹配")
		}
	}

}
func showSmsTypeMenu() {
	var (
		key     int
		content string
		userid  int
	)

	for {
		fmt.Println("----------欢迎使用聊天------------")
		fmt.Println("\t\t\t1.世界广播")
		fmt.Println("\t\t\t2.偷偷私发")
		fmt.Println("\t\t\t3.退出聊天室")
		fmt.Println("\t\t\t请选择<1~3>")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			{
				fmt.Println("请输入要发送的语言")
				fmt.Scanf("%s\n", &content)
				sp := &SmsProcess{}
				err := sp.sendGroupSes(content)
				if err != nil {
					fmt.Println("世界广播失败,请重试")
				}
			}

		case 2:
			{
				fmt.Println("请输入要发送的语言")
				fmt.Scanf("%s\n", &content)
				fmt.Println("请输入要发送给的用户id")
				fmt.Scanf("%d\n", &userid)
				//处理逻辑
				sp := &SmsProcess{}
				err := sp.sendPrivateMes(content, userid)
				if err != nil {
					fmt.Println("偷偷私发,请重试")
				}
			}
		case 3:
			{
				goto loop
			}

		}

	}
loop:
}
func ExitOnline(userid int, tf *comutils.Transfer, name string, pwd string) {
	var mes message.Message
	mes.Type = message.DownOnlineMesType
	var exitmes message.DownOnlineMes
	exitmes.Userid = userid
	exitmes.Username = name
	exitmes.Userpwd = pwd
	exitmes.UserStatus = message.UserOffline
	data, err := json.Marshal(exitmes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}
	//发消息
	err = tf.PakWrite(data)
	if err != nil {
		return
	}
}
