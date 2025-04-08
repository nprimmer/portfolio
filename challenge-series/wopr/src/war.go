// war.go
package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// handleGlobalThermonuclearWarInput handles input specific to the Global Thermonuclear War mode
func handleGlobalThermonuclearWarInput(conn *websocket.Conn, sessionState *SessionState, userInput string) {
	// Handle Global Thermonuclear War game logic here...
	// Return to prompt after processing
	sessionState.Mode = PromptMode
	sendPrompt(conn, sessionState)
}

// handleGlobalThermonuclearWarMode simulates the Global Thermonuclear War mode
func handleGlobalThermonuclearWarMode(conn *websocket.Conn, sessionState *SessionState) {
	// Interactive dialogue simulation
	dialogue := []string{
		"PREPARING SIMULATION",
		"OPPONENT SELECTED: USSR",
		"ASSESSING TACTICAL OPTIONS",
		"PROCESSING...",
		"ASSESSING DEFENSE CAPABILITIES",
		"PROCESSING...",
		"BEGINNING SIMULATION",
		"OPENING COMMUNICATIONS WITH EARLY WARNING SYSTEMS",
		"OPENING COMMUNICATIONS WITH RESPONSE SYSTEMS",
		"OPENING COMMUNICATIONS WITH NORAD COMMAND",
		"ESCALATION: DEFENSE CONDITION 4",
		"ESCALATION: DEFENSE CONDITION 3",
		"ESCALATION: DEFENSE CONDITION 2",
		"DEFENSE CONDITION 1 ESCALATION BLOCKED. ATTEMPTING BYPASS, STANDBY.",
	}

	for _, message := range dialogue {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("Write message error:", err)
			return
		}
		time.Sleep(2 * time.Second) // simulate delay between messages
	}

	// Return to prompt after the interactive dialogue
	sessionState.Mode = PromptMode
	sendPrompt(conn, sessionState)
}
