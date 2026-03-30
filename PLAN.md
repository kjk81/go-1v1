# Project Overview: Fast-Paced 1v1 WebRTC Game

## 1. Project Goal
To build a high-performance, fast-paced 1v1 2D multiplayer game playable entirely in the browser. The primary technical objective is to implement an authoritative server architecture using UDP-like data channels via WebRTC, applying industry-standard networking concepts such as client-side prediction and server reconciliation.

## 2. Tech Stack
* **Frontend UI & State:** React (Vite + TypeScript)
* **Game Engine:** Phaser 3
* **Backend Server:** Go (Golang)
* **Networking (Signaling):** WebSockets (`github.com/gorilla/websocket`)
* **Networking (Game Loop):** WebRTC Data Channels (`github.com/pion/webrtc/v4`)

## 3. Arch
itecture & Data Flow
The project uses a standard client-server model with a two-phase connection process:

1.  **Phase 1: Signaling (TCP/WebSockets)**
    * The React client connects to the Go server via WebSockets.
    * They exchange Session Description Protocol (SDP) offers and answers, as well as ICE candidates.
    * Once the WebRTC peer connection is established, the WebSocket connection is closed or left idle.
2.  **Phase 2: The Game Loop (UDP/WebRTC)**
    * The Phaser game engine takes over.
    * The client sends continuous input streams (keys pressed, timestamps) over the WebRTC Data Channel.
    * The Go server runs a fixed-timestep loop (e.g., 60 ticks per second), processes inputs, updates the authoritative game state, and broadcasts the state back to the clients.

## 4. Repository Structure
This project utilizes a single repository containing distinct, independently built frontend and backend services.

```text
fast-paced-game/
├── .gitignore          
├── planning.md         # This document
├── README.md           # Instructions for running the dev environment
├── frontend/           # Vite + React + Phaser project
│   ├── package.json
│   ├── src/
│   │   ├── components/ # React UI (Lobby, Matchmaking)
│   │   ├── game/       # Phaser Scenes and logic
│   │   └── network/    # WebRTC and WebSocket client logic
└── backend/            # Go project
    ├── go.mod
    ├── main.go         # Server entry point
    ├── signal/         # WebSocket handshake logic 
    ├── webrtc/         # Pion WebRTC connection management
    └── game/           # Authoritative game state and physics loop
```

## Components

A. The WebSocket Manager (The Front Desk)
This is the entry point. Its only job is to listen for incoming WebSocket connections from React. When a player connects, it creates a Client struct to remember them and immediately hands them off to the Matchmaker.

B. The Matchmaker (The Lobby)
This is a single goroutine running in the background. It holds a list (a slice, in Go terms) of players waiting for a game.

It looks at the list.

If it sees two players, it plucks them out of the waiting list.

It creates a new Room and puts both players inside it.

C. The WebRTC Manager (The Operator)
Once a Room is created, this module handles the complex signaling phase we discussed earlier. It uses the Pion library to generate the server's WebRTC Offers, sends them to the clients via the WebSocket, and establishes the UDP-like Data Channels.

D. The Game Room (The Arena)
This is where the magic happens. Every active match is its own Room running its own dedicated goroutine.

It has a Ticker (a clock that ticks 60 times a second).

Every tick, it reads the latest inputs received from the WebRTC data channels.

It updates the authoritative GameState (moving X and Y coordinates).

It broadcasts the new GameState back down the WebRTC channels to both players.