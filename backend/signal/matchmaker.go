package signal

type MatchMaker struct {
	Mail          chan Client
	waitingPlayer Client
	isWaiting     bool
}

func (m MatchMaker) Run() {
	for p := range m.Mail {
		if m.isWaiting {
			// pass players in room to signaling
			r := Room{C1: m.waitingPlayer, C2: p}
			go r.Start()
			m.isWaiting = false
		} else {
			m.waitingPlayer = p
		}
	}
}
