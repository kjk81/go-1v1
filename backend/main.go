package main

import (
	"net/http"

	"fmt"

	"game-server/signal"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

var mm *signal.MatchMaker

func main() {
	http.HandleFunc("/", handleJoin) // set route
	mm = new(signal.MatchMaker)
	go mm.Run() // matchmaking subroutine
}

func handleJoin(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // upgrade request to websocket
	if err != nil {
		fmt.Println(err)
	}
	// add Player struct to matchmaker to process
	player := signal.Client{
		Socket: conn,
	}
	mm.Mail <- player
}

// validating request origin is allowed
func checkOrigin(r *http.Request) bool {
	return true // temp for dev
}
