package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"greatestworks/network"
	"greatestworks/network/protocol/gen/player"
	"strconv"
)

type MessageHandler func(message *network.ClientPacket)

type InputHandler func(param *InputParma)

// CreatePlayer 创建角色
func (c *Client) CreatePlayer(param *InputParma) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 2 {
		return
	}
	msg := &player.CSCreatePlayer{
		UserName: param.Param[0],
		Password: param.Param[1],
	}
	c.Transport(id, msg)
}

func (c *Client) OnCreatePlayerRsp(packet *network.ClientPacket) {
	fmt.Println("恭喜创建角色成功！")
}

// Login 登入
func (c *Client) Login(param *InputParma) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 2 {
		fmt.Println("输入参数数据有误！")
		return
	}
	msg := &player.CSLogin{
		UserName: param.Param[0],
		Password: param.Param[1],
	}
	c.Transport(id, msg)
}

func (c *Client) OnLoginRsp(packet *network.ClientPacket) {
	rsp := &player.SCLogin{}
	err := proto.Unmarshal(packet.Msg.Data, rsp)
	if err != nil {
		return
	}
	fmt.Printf("登入成功")
}

func (c *Client) AddFriend(param *InputParma) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	msg := &player.CSAddFriend{
		UId: parseUint,
	}

	c.Transport(id, msg)
}

func (c *Client) OnAddFriendRsp(packet *network.ClientPacket) {
	rsp := &player.SCADDFriend{}
	err := proto.Unmarshal(packet.Msg.Data, rsp)
	if err != nil {
		return
	}
	fmt.Println("添加好友成功！")
}

func (c *Client) DelFriend(param *InputParma) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	msg := &player.CSDelFriend{
		UId: parseUint,
	}

	c.Transport(id, msg)
}
func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {
	//rsp := &player.SCDelFriend{}
	//err := proto.Unmarshal(packet.Msg.Data, rsp)
	//if err != nil {
	//	return
	//}
	fmt.Println("删除好友成功！")
}

func (c *Client) SendChatMsg(param *InputParma) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 3 || len(param.Param[0]) == 0 {
		return
	}

	parseUId, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	parseCategory, err := strconv.ParseInt(param.Param[2], 10, 32)
	if err != nil {
		return
	}

	msg := &player.CSSendChatMsg{
		UId: parseUId,
		Msg: &player.ChatMessage{
			Content: param.Param[1],
			Extra:   nil,
		},
		Category: int32(parseCategory),
	}

	c.Transport(id, msg)
}
func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {
	//rsp := &player.SCSendChatMsg{}
	//err := proto.Unmarshal(packet.Msg.Data, rsp)
	//if err != nil {
	//	return
	//}
	fmt.Println("发送消息成功！")
}
