package signal

import "game-server/game"

type Room struct {
	C1   Client
	C2   Client
	Game game.Game
}

func (r Room) Start() {
	StartClientPeer(r.C1)
	StartClientPeer(r.C2)
}
