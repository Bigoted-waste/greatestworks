package main

import (
	"github.com/phuhao00/sugar"
	"greatestworks/world"
)

func main() {
	world.MM = world.NewMgrMgr()
	go world.MM.Pm.Run()
	sugar.WaitSignal(world.MM.OnSystemSignal)
}
