package processes

import (
	"encoding/json"
	"fmt"
	"mysqlandcon/common/comutils"
	"mysqlandcon/common/message"
	"mysqlandcon/sever/model"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//定义连接属于哪个用户
	Userid int
}
type ExitOnline struct {
}

// 通知其他人我上线了
func (this *UserProcess) NotifyOthers(userid int, stastus int) {
	for id, up := range Usermgr.UserOnline {
		//过滤掉自己
		if id == userid {
			continue
		}
		//新方法--up是其他用户，userid是新登录的用户id
		up.Notify(userid, stastus)
	}
}
func (this *UserProcess) Notify(userid int, stastus int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyuserstatusmes message.NotifyUserStatusMes
	notifyuserstatusmes.Userid = userid
	notifyuserstatusmes.Status = stastus
	//notifyuserstatusmes进行序列化
	data, err := json.Marshal(notifyuserstatusmes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	//对mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}
	//发送消息给客户端
	tf := &comutils.Transfer{Conn: this.Conn}
	err = tf.PakWrite(data)
	if err != nil {
		return
	}
}
func (this *UserProcess) SeverProcessLogin(mes *message.Message) (err error) {
	//先从mes中取出data，将其反序列化
	var LogInMes message.LogInMes
	err = json.Unmarshal([]byte(mes.Data), &LogInMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return err
	}
	//声明返回变量，并且序列化
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var LoginResMes message.LoginResMes
	//使用redis进行验证
	user, err := model.MyuserDao.Login(LogInMes.UserId, LogInMes.UserPwd, LogInMes.Username)
	if err != nil {
		if err == model.ERR_USER_NOTEX {
			LoginResMes.Code = 500
			LoginResMes.Err = err.Error()
		} else if err == model.ERR_USER_PWD {
			LoginResMes.Code = 403
			LoginResMes.Err = err.Error()
		} else if err == model.ERR_USERNAME {
			LoginResMes.Code = 600
			LoginResMes.Err = err.Error()
		} else if err == model.ERR_USER_ONLINED {
			LoginResMes.Code = 700
			LoginResMes.Err = err.Error()
		} else {
			LoginResMes.Err = "服务器内部错误..."
		}
	} else {
		LoginResMes.Code = 200
		this.Userid = LogInMes.UserId
		Usermgr.UpdateOnlineUser(this, message.UserOnline)
		//通知其他用户我上线了
		this.NotifyOthers(LogInMes.UserId, message.UserOnline)
		//将在线用户id返回给客户端
		for id, _ := range Usermgr.UserOnline {
			LoginResMes.UsersId = append(LoginResMes.UsersId, id)
		}
		fmt.Println(user, "登录ok")
	}

	data, err := json.Marshal(LoginResMes)
	if err != nil {
		return err
	}
	resMes.Data = string(data)
	//准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		return err
	}
	tf := &comutils.Transfer{Conn: this.Conn}
	err = tf.PakWrite(data)
	if err != nil {
		return err
	}
	return nil
}
func (this *UserProcess) SeverProcessSignup(mes *message.Message) (err error) {
	//先从mes中取出data，将其反序列化
	var SignUpMes message.SignupMes
	err = json.Unmarshal([]byte(mes.Data), &SignUpMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return err
	}
	//声明返回变量，并且序列化
	var resMes message.Message
	resMes.Type = message.SignupResMesType
	var SignupResMes message.SignupResMes
	//使用redis进行验证
	err = model.MyuserDao.Signup(SignUpMes.UserId, SignUpMes.UserPwd, SignUpMes.UserName)
	if err != nil {
		if err == model.ERR_USER_EX {
			SignupResMes.Code = 400
			SignupResMes.Err = err.Error()
		} else {
			SignupResMes.Err = "服务器内部错误..."
		}
	} else {
		SignupResMes.Code = 200
		fmt.Println("注册ok")
	}
	data, err := json.Marshal(SignupResMes)
	if err != nil {
		return err
	}
	resMes.Data = string(data)
	//准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		return err
	}
	tf := &comutils.Transfer{Conn: this.Conn}
	err = tf.PakWrite(data)
	if err != nil {
		return err
	}
	return nil
}
func (eo *ExitOnline) ProcessExitOnline(mes *message.Message) {
	var exitonline message.DownOnlineMes
	err := json.Unmarshal([]byte(mes.Data), &exitonline)
	if err != nil {
		return
	}
	for id, up := range Usermgr.UserOnline {
		if id == exitonline.Userid {
			delete(Usermgr.UserOnline, id)
			model.MyuserDao.ExitOnline(exitonline.Userid, exitonline.Userpwd, exitonline.Username)
			continue
		}
		data, err := json.Marshal(*mes)
		if err != nil {
			return
		}
		tf := &comutils.Transfer{Conn: up.Conn}
		err = tf.PakWrite(data)
		if err != nil {
			return
		}

	}
}
