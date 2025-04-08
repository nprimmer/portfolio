// password.go
package main

import (
	"github.com/gorilla/websocket"
)

func handlePasswordInput(conn *websocket.Conn, sessionState *SessionState, userInput string) {
	if userInput == invalidPasswordMessage {
		conn.WriteMessage(websocket.TextMessage, []byte("INVALID PASSWORD"))
		conn.WriteMessage(websocket.TextMessage, []byte(invalidPasswordResponse))
		conn.WriteMessage(websocket.TextMessage, []byte("ENTER PASSWORD:"))
	} else if message, valid := validPasswords[userInput]; valid {
		if sessionState.GlobalThermonuclearWar {
			conn.WriteMessage(websocket.TextMessage, []byte("THE SIMULATION IS ALREADY ENGAGED"))
			sessionState.Mode = PromptMode
			sendPrompt(conn, sessionState)
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte(message))
			conn.WriteMessage(websocket.TextMessage, []byte("ACCESS GRANTED. YOU MAY PROCEED."))
			conn.WriteMessage(websocket.TextMessage, []byte("ACCESS ID: GC24{6C3E3B4F-31BE-405D-8793-6238DE2695A4}"))

			// Set the session state
			sessionState.GlobalThermonuclearWar = true
			sessionState.Mode = GlobalThermonuclearWarMode

			// Handle the interactive dialogue
			handleGlobalThermonuclearWarMode(conn, sessionState)
		}
	} else {
		conn.WriteMessage(websocket.TextMessage, []byte("INVALID PASSWORD"))
		conn.WriteMessage(websocket.TextMessage, []byte("ENTER PASSWORD:"))
	}
}
