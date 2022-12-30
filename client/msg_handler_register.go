package main

import "greatestworks/network/protocol/gen/messageId"

func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[messageId.MessageId_SCLogin] = c.OnLoginRsp
	c.messageHandlers[messageId.MessageId_SCADDFriend] = c.OnAddFriendRsp
	c.messageHandlers[messageId.MessageId_SCDelFriend] = c.OnDelFriendRsp
	c.messageHandlers[messageId.MessageId_SCSendChatMsg] = c.OnSendChatMsgRsp
}
