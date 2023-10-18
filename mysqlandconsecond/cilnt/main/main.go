package main

import (
	"fmt"
	"mysqlandcon/cilnt/process"
	"os"
)

func main() {
	var (
		key    int
		userid int
		pwd    string
		name   string
	)

loop:
	for {
		fmt.Printf("-------------------欢迎登录多人聊天系统------------------\n")
		fmt.Printf("\t\t\t1.登录多人聊天系统\n")
		fmt.Printf("\t\t\t2.注册账号\n")
		fmt.Printf("\t\t\t3.退出系统\n")
		fmt.Printf("\t\t\t请输入所选<1~3>\n")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			{
				fmt.Printf("----------------正在登录中---------\n")
				fmt.Printf("请输入用户id\n")
				fmt.Scanf("%d\n", &userid)
				fmt.Printf("请输入用户名密码\n")
				fmt.Scanf("%s\n", &pwd)
				fmt.Printf("请输入用户名\n")
				fmt.Scanf("%s\n", &name)
				up := &process.UserProcess{}
				err := up.Login(userid, pwd, name)
				if err != nil {
					fmt.Println(err)
				}
			}
		case 2:
			{
				fmt.Printf("-----------------正在注册中---------\n")
				fmt.Printf("请输入用户名id\n")
				fmt.Scanf("%d\n", &userid)
				fmt.Printf("请输入用户名密码\n")
				fmt.Scanf("%s\n", &pwd)
				fmt.Printf("请输入用户名\n")
				fmt.Scanf("%s\n", &name)
				up := &process.UserProcess{}
				err := up.Signup(userid, pwd, name)
				if err != nil {
					fmt.Println(err)
				}
				goto loop
			}
		case 3:
			{
				fmt.Printf("已退出系统\n")
				os.Exit(0)

			}
		default:
			fmt.Printf("输入有误，请重新输入\n")
		}

	}
}
