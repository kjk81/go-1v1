package signal

import (
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

var peerConfig = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs: []string{"stun:stun.l.google.com:19302"},
		},
	},
}

type Client struct {
	RTC    *webrtc.PeerConnection
	Socket *websocket.Conn
}

// uses gorilla socket to connect peers
func StartClientPeer(c Client) {
	conn, err := webrtc.NewPeerConnection(peerConfig)
	if err != nil {
		panic(err)
	}
	c.RTC = conn
}
