package network

import "net"

// SessionPacket 会话数据包
type SessionPacket struct {
	Msg  *Message
	Sess *Session
}

type ClientPacket struct {
	Msg  *Message
	Conn net.Conn
}
