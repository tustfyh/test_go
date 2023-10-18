package processes

import (
	"fmt"
	"mysqlandcon/common/message"
)

var (
	Usermgr *UserMgr
)

type UserMgr struct {
	UserOnline map[int]*UserProcess //拿到每个用户的net.conn
}

// 对usermgr进行初始化
// 在主程序运行前会自动初始化
func init() {
	Usermgr = &UserMgr{make(map[int]*UserProcess, 1024)}
}

// 对usermar的更新操作
func (this *UserMgr) UpdateOnlineUser(up *UserProcess, status int) {
	if status == message.UserOnline {
		Usermgr.UserOnline[up.Userid] = up
	} else {
		this.DelOnlineUser(up.Userid)
	}
}

// del操作
func (this *UserMgr) DelOnlineUser(userid int) {
	delete(this.UserOnline, userid)
}

// 返回当前在线用户列表
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.UserOnline
}

// 根据id返回对应的值
func (this *UserMgr) GetAllOnlineUserById(userid int) (up *UserProcess, err error) {
	up, ok := this.UserOnline[userid]
	//未找到处理
	if ok == false {
		err = fmt.Errorf("用户%d不在线或不存在", userid)
		return nil, err
	}
	return up, nil
}
