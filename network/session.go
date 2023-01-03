package network

import (
	"encoding/binary"
	"fmt"
	"greatestworks/log"
	"net"
	"time"
)

type Session struct {
	UId            uint64
	conn           net.Conn
	IsClose        bool
	packer         *NormalPacker
	WriteCh        chan *Message
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
	//
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian),
		WriteCh: make(chan *Message, 1),
	}
}

func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

func (s *Session) Read() {
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	for {
		message, err := s.packer.Unpack(s.conn)
		if err != nil {
			log.Logger.ErrorF(err.Error())
			continue
		}
		log.Logger.InfoF("服务端接受消息: ", string(message.Data))
		s.WriteCh <- &Message{
			Id:   99,
			Data: []byte("pong"),
		}
	}
}

func (s *Session) Write() {
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Logger.ErrorF(err.Error())
		return
	}
	for {
		select {
		case resp := <-s.WriteCh:
			s.send(resp)
		}
	}
}

func (s *Session) send(message *Message) {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Logger.ErrorF(err.Error())
		return
	}
	bytes, err := s.packer.Pack(message)
	if err != nil {
		return
	}
	_, err = s.conn.Write(bytes)
	if err != nil {
		log.Logger.ErrorF(err.Error())
	}
}

func (s *Session) SendMsg(msg *Message) {
	s.WriteCh <- msg
}
