package player

import (
	"google.golang.org/protobuf/proto"
	"greatestworks/function"
	"greatestworks/log"
	"greatestworks/network"
	"greatestworks/network/protocol/gen/player"
)

type Handler func(msg *network.Message)

func (p *Player) AddFriend(msg *network.Message) {
	req := &player.CSAddFriend{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	if !function.CheckInNumberSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

func (p *Player) DelFriend(msg *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	p.FriendList = function.DelEleInSlice(req.UId, p.FriendList)
}

func (p *Player) ResolveChatMsg(msg *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	log.Logger.InfoF(req.Msg.Content)
	// todo 收到消息 然后转发给客户端（当你的好友非你发消息情况）
}
