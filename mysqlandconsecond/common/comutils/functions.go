package comutils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"mysqlandcon/common/message"
	"net"
)

// 传输者
type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (this *Transfer) PakWrite(data []byte) (err error) {
	var buf [4]byte
	var paklen uint32
	paklen = uint32(len(data))
	binary.BigEndian.PutUint32(buf[:], paklen)
	n, err := this.Conn.Write(buf[:])
	//fmt.Println(buf)
	if n != 4 || err != nil {
		return err
	}
	//发送消息本身
	_, err = this.Conn.Write(data)
	//fmt.Printf("%s\n", string(data))
	if err != nil {
		return err
	}
	return nil
}
func (tf *Transfer) PakRed() (msg message.Message, err error) {
	n, err := tf.Conn.Read(tf.Buf[0:4])
	if n != 4 {
		return message.Message{}, err
	}
	var packLen uint32
	packLen = binary.BigEndian.Uint32(tf.Buf[0:4])
	n, err = tf.Conn.Read(tf.Buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	err = json.Unmarshal(tf.Buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}
