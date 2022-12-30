package world

import "greatestworks/network/protocol/gen/messageId"

func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[messageId.MessageId_CSLogin] = mm.UserLogin
	mm.Handlers[messageId.MessageId_CSCreatePlayer] = mm.CreatePlayer
}
