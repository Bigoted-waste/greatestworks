package manager

import (
	"greatestworks/business/player"
)

// PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players map[uint64]*player.Player
	addPCh  chan player.Player
}

// Add ...
func (pm *PlayerMgr) Add(p *player.Player) {
	pm.players[p.Uid] = p
	go p.Run()
}

// Del ...
func (pm *PlayerMgr) Del(p player.Player) {
	delete(pm.players, p.Uid)
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(&p)
		}
	}
}

func (pm *PlayerMgr) GetPlayer(uId uint64) *player.Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
