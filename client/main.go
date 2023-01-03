package main

import "github.com/phuhao00/sugar"

func main() {
	c := NewClient()
	// 添加应用处理器
	c.InputHandlerRegister()
	// 添加消息处理器
	c.MessageHandlerRegister()
	c.Run()
	sugar.WaitSignal(c.OnSystemSignal)

	//select {} //阻塞
}
