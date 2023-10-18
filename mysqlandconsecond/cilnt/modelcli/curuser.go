package modelcli

import (
	"mysqlandcon/sever/model"
	"net"
)

type CurUser struct {
	model.User
	Conn net.Conn
}
