package player

import (
	"github.com/phuhao00/sugar"
	"google.golang.org/protobuf/proto"
	"greatestworks/aop/logger"
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
	if !sugar.CheckInSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

func (p *Player) DelFriend(msg *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	p.FriendList = sugar.DelOneInSlice(req.UId, p.FriendList)
}

func (p *Player) ResolveChatMsg(msg *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(msg.Data, req)
	if err != nil {
		return
	}
	logger.Logger.InfoF(req.Msg.Content)
	// todo 收到消息 然后转发给客户端（当你的好友非你发消息情况）
}
