package main

import "fmt"

type Account struct {
	balance      float64
	addnum       []float64
	subnum       []float64
	reasonforadd []string
	reasonforsub []string
}

func (account *Account) Addandrea(addnum float64, reason string) {
	account.balance += addnum
	account.addnum = append(account.addnum, addnum)
	account.reasonforadd = append(account.reasonforadd, reason)
}
func (account *Account) Subandrea(subnum float64, reason string) {
	account.balance -= subnum
	account.subnum = append(account.addnum, subnum)
	account.reasonforsub = append(account.reasonforsub, reason)
}

// 简单的go小项目---家庭收支明细
func main() {
	var (
		key    int
		count  float64
		reason string
	)
	account := Account{
		balance:      10000,
		addnum:       nil,
		subnum:       nil,
		reasonforadd: nil,
		reasonforsub: nil,
	}
	account.reasonforadd = make([]string, 0)
	account.reasonforsub = make([]string, 0)
	account.addnum = make([]float64, 0)
	account.subnum = make([]float64, 0)
	loop := true
	for loop {
		fmt.Println("----------家庭收支明细软件----------")
		fmt.Println("----------1.收支明细查看-----------")
		fmt.Println("----------2.登记收入---------------")
		fmt.Println("----------3.登记支出---------------")
		fmt.Println("----------4.退出软件---------------")
		fmt.Printf("请选择（1~4）")
		fmt.Scan(&key)
		switch key {
		case 1:
			{
				fmt.Println("----------当前收支明细-----------------")
				fmt.Printf("当前余额%v元\n", account.balance)
				for i := 0; i < len(account.addnum); i++ {
					fmt.Printf("收入%v元,原因为%v\n", account.addnum[i], account.reasonforadd[i])
				}
				for i := 0; i < len(account.subnum); i++ {
					fmt.Printf("登出%v元,原因为%v\n", account.subnum[i], account.reasonforsub[i])
				}
				fmt.Printf("------------------------------------\n")
			}

		case 2:
			{
				fmt.Printf("请输入收入金额与收入金额原因\n")
				fmt.Scan(&count, &reason)
				account.Addandrea(count, reason)
			}
		case 3:
			{
				fmt.Printf("请输入登出金额与登出金额的原因\n")
				fmt.Scan(&count, &reason)
				account.Subandrea(count, reason)
			}
		case 4:
			{
				loop = false
			}

		default:
			{
				fmt.Println("输入数字有误，请重新输入")
			}

		}
	}
	fmt.Println("你已经退出该软件，感谢使用")
}
