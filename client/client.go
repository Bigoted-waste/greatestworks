package main

import "C"
import (
	"fmt"
	"greatestworks/network"
	"greatestworks/network/protocol/gen/messageId"
	"os"
	"syscall"
)

type Client struct {
	cli             *network.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[messageId.MessageId]MessageHandler
	console         *ClientConsole
	chInput         chan *InputParma
}

func NewClient() *Client {
	c := &Client{
		cli:             network.NewClient(":8023"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[messageId.MessageId]MessageHandler{},
		console:         NewClientConsole(),
	}
	c.cli.OnMessage = c.OnMessage
	c.cli.ChMsg = make(chan *network.Message, 1)
	c.chInput = make(chan *InputParma, 1)
	c.console.chInput = c.chInput
	return c
}

func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput:
				fmt.Printf("cmd:%s,param:%v \t \n", input.Command, input.Param)
				inputHandler := c.inputHandlers[input.Command]
				if inputHandler != nil {
					inputHandler(input)
				}
			}
		}
	}()
	// 启动控制台
	go c.console.Run()
	// 启动客户端
	go c.cli.Run()
}

func (c *Client) OnMessage(packet *network.ClientPacket) {
	if handler, ok := c.messageHandlers[messageId.MessageId(packet.Msg.Id)]; ok {
		handler(packet)
	}
}

func (c *Client) OnSystemSignal(signal os.Signal) bool {
	fmt.Printf("[client] 收到信号 %v \n", signal)
	tag := true
	switch signal {
	case syscall.SIGHUP:
		// TODO
	case syscall.SIGPIPE:
		// TODO
	default:
		fmt.Println("[Client] 收到信号准备退出。。。")
		tag = false
	}
	return tag
}
