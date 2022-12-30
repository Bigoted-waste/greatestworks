package main

import (
	"google.golang.org/protobuf/proto"
	"greatestworks/network"
	"greatestworks/network/protocol/gen/messageId"
)

// InputHandlerRegister 注册输出handler
func (c *Client) InputHandlerRegister() {
	c.inputHandlers[messageId.MessageId_CSLogin.String()] = c.Login
	c.inputHandlers[messageId.MessageId_CSADDFriend.String()] = c.AddFriend
	c.inputHandlers[messageId.MessageId_CSDelFriend.String()] = c.DelFriend
	c.inputHandlers[messageId.MessageId_CSSendChatMsg.String()] = c.SendChatMsg
}

// GetMessageIdByCmd 通过Handler获取messageId
func (c *Client) GetMessageIdByCmd(cmd string) messageId.MessageId {
	mid, ok := messageId.MessageId_value[cmd]
	if ok {
		return messageId.MessageId(mid)
	}
	return messageId.MessageId_None
}

func (c *Client) Transport(id messageId.MessageId, msg proto.Message) {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	c.cli.ChMsg <- &network.Message{
		Id:   uint64(id),
		Data: bytes,
	}
}
