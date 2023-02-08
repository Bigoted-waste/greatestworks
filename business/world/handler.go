package world

import (
	"google.golang.org/protobuf/proto"
	"greatestworks/aop/logger"
	logicPlayer "greatestworks/business/player"
	"greatestworks/network"
	"greatestworks/network/protocol/gen/player"
	"time"
)

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	msg := &player.CSLogin{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	newPlayer := logicPlayer.NewPlayer()
	newPlayer.Uid = uint64(time.Now().Unix())
	newPlayer.HandlerParamCh = message.Sess.WriteCh
	message.Sess.IsPlayerOnline = true
	mm.Pm.Add(newPlayer)
	newPlayer.Run()
}

func (mm *MgrMgr) CreatePlayer(message *network.SessionPacket) {
	msg := &player.CSCreatePlayer{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil { //
		return
	}
	logger.Logger.InfoF("[MgrMgr.CreatePlayer] >>>>", msg)
	mm.SendMsg(message.Msg.Id, &player.SCCreatePlayer{}, message.Sess)
}

func (mm *MgrMgr) SendMsg(id uint64, message proto.Message, session *network.Session) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	rsp := &network.Message{
		Id:   id,
		Data: bytes,
	}
	session.SendMsg(rsp)
}
