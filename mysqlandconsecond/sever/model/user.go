package model

type User struct {
	//必须加tag，要不然和数据库字段对不上
	UserId     int    `json:"userid"`
	UserName   string `json:"username"`
	UserPwd    string `json:"userpwd"`
	UserStatus int    `json:"userstatus"`
}
