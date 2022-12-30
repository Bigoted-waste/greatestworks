package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClientConsole struct {
	chInput chan *InputParma
}

type InputParma struct {
	Command string
	Param   []string
}

func NewClientConsole() *ClientConsole {
	c := &ClientConsole{}
	return c
}

func (c ClientConsole) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input err,check your input and try again !!!")
			continue
		}
		readString = strings.Replace(readString, "\n", "", -1)
		split := strings.Split(readString, " ")
		if len(split) == 0 {
			fmt.Println("input err,check your input and try again !!!")
			continue
		}
		in := &InputParma{
			Command: split[0],
			Param:   split[1:],
		}
		c.chInput <- in
	}
}
