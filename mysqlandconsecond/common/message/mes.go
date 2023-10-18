package message

const (
	LogInMesType            = "LogInMes"
	LoginResMesType         = "LoginResMes"
	SignupMesType           = "SignupMes"
	SignupResMesType        = "SignupResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SmsPrivateMesType       = "SmsPrivateMes"
	DownOnlineMesType       = "DownOnlineMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
type LogInMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	Username string `json:"username"`
}
type LoginResMes struct {
	Code    uint   `json:"code"` // 200表示登录成功，500表示该用户不存在，403表示密码错误，600表示用户名错误
	UsersId []int  `json:"usersId"`
	Err     string `json:"err"`
}
type SignupMes struct { //处理注册的逻辑结构体
	UserId     int    `json:"userid"`
	UserName   string `json:"username"`
	UserPwd    string `json:"userpwd"`
	UserStatus int    `json:"userstatus"`
}
type SignupResMes struct {
	Code uint   `json:"code"` //400表示该用户id已被注册，200表示成功
	Err  string `json:"err"`
}
type NotifyUserStatusMes struct {
	Userid int `json:"userid"`
	Status int `json:"status"`
}
type SmsMes struct {

	//必须加tag，要不然和数据库字段对不上
	UserId     int    `json:"userid"`
	UserName   string `json:"username"`
	UserPwd    string `json:"userpwd"`
	UserStatus int    `json:"userstatus"`
	Content    string `json:"content"` //发送的消息
}
type SmsPrivateMes struct {
	Content       string `json:"content" `
	Userid        int    `json:"userid" `
	Username      string `json:"username"`
	PrivateUserid int    `json:"privateUserid"`
	UserStatus    int    `json:"userStatus"`
}
type DownOnlineMes struct {
	Userid     int    `json:"userid"`
	Username   string `json:"username"`
	Userpwd    string `json:"userpwd"`
	UserStatus int    `json:"userStatus"`
}

const (
	UserOnline  = 1
	UserOffline = 0
)
