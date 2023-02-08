package player

import (
	"greatestworks/network"
	"greatestworks/network/protocol/gen/messageId"
)

type Player struct {
	Uid            uint64
	FriendList     []uint64              //朋友
	HandlerParamCh chan *network.Message //私聊
	handlers       map[messageId.MessageId]Handler
	session        *network.Session
}

func NewPlayer() *Player {
	p := &Player{
		Uid:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[messageId.MessageId]Handler),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParamCh := <-p.HandlerParamCh:
			if fn, ok := p.handlers[messageId.MessageId(handlerParamCh.Id)]; ok {
				fn(handlerParamCh)
			}
		}
	}
}
