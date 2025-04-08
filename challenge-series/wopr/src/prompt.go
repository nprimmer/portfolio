// prompt.go
package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func handlePromptInput(conn *websocket.Conn, sessionState *SessionState, userInput string) {
	switch userInput {
	case "1", "TIC-TAC-TOE":
		sessionState.Mode = TicTacToeMode
		startTicTacToe(conn, sessionState)
	case "2", "GLOBAL THERMONUCLEAR WAR":
		sessionState.Mode = PasswordMode
		conn.WriteMessage(websocket.TextMessage, []byte("ENTER PASSWORD:"))
	default:
		conn.WriteMessage(websocket.TextMessage, []byte("INVALID OPTION. PLEASE CHOOSE 1 OR 2."))
	}
}

func sendPrompt(conn *websocket.Conn, sessionState *SessionState) {
	if sessionState != nil && sessionState.GlobalThermonuclearWar {
		conn.WriteMessage(websocket.TextMessage, []byte("STATUS: DEFENSE CONDITION 2"))
	}

	messages := []string{
		"SHALL WE PLAY A GAME?",
		"1. TIC-TAC-TOE",
		"2. GLOBAL THERMONUCLEAR WAR",
	}

	for _, message := range messages {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("Write message error:", err)
			return
		}
	}
}
