// main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Mode represents the current mode of the application
type Mode int

const (
	PromptMode Mode = iota
	TicTacToeMode
	GlobalThermonuclearWarMode
	PasswordMode
)

// SessionState holds the state of a session
type SessionState struct {
	ID                     string
	GlobalThermonuclearWar bool
	Mode                   Mode
	TicTacToeBoard         [9]string
	CurrentPlayer          string
	ComputerWins           int
	PlayerWins             int
	Draws                  int
}

var (
	upgrader       = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	sessions       = make(map[string]*SessionState)
	sessionsMu     sync.Mutex
	validPasswords = map[string]string{
		"LOVELACE": "ACCESS RECORD 14A8A56B-: PASSWORD ACCEPTED. WELCOME DR WINTERS.",
		"THOMPSON": "ACCESS RECORD EA93-4282-B441: PASSWORD ACCEPTED. WELCOME DR WINTERS.",
		"ANSEL":    "ACCESS RECORD -2170E820A5AE: PASSWORD ACCEPTED. WELCOME DR WINTERS.",
	}
	invalidPasswordMessage  = "JOSHUA"
	invalidPasswordResponse = "INVALID PASSWORD. INCIDENT LOG GC24{701287BF-5D3F-4FE8-B6F3-9C25C92E9284}"
)

// generateSessionID generates a random session ID
func generateSessionID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%08x", rand.Int63())
}

// handleNewConnection handles a new WebSocket connection
func handleNewConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Generate and store session ID
	sessionID := generateSessionID()
	sessionState := &SessionState{ID: sessionID, Mode: PromptMode}
	sessionsMu.Lock()
	sessions[sessionID] = sessionState
	sessionsMu.Unlock()

	// Initial connection messages
	initialMessages := []string{
		"WAR OPERATION PLAN RESPONSE",
		"DEVELOPED 1983 - ALEX WINTERS",
		"SESSION ID: " + sessionID,
		"TO RESUME THIS SESSION, CONNECT TO /WOPR/" + sessionID,
	}
	for _, message := range initialMessages {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("Write message error:", err)
			return
		}
	}

	// Send prompt
	sendPrompt(conn, sessionState)

	// Handle user input
	handleUserInput(conn, sessionState)
}

// handleResumeConnection handles resuming a WebSocket connection with a given session ID
func handleResumeConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	sessionsMu.Lock()
	sessionState, exists := sessions[sessionID]
	sessionsMu.Unlock()

	if !exists {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Resume connection messages
	resumeMessages := []string{
		"RESUMING SESSION...",
		"SESSION ID: " + sessionID,
	}
	for _, message := range resumeMessages {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("Write message error:", err)
			return
		}
	}

	// Send prompt or resume current mode
	sendPrompt(conn, sessionState)

	// Handle user input
	handleUserInput(conn, sessionState)
}

// handleUserInput handles the user input for both new and resumed connections
func handleUserInput(conn *websocket.Conn, sessionState *SessionState) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read message error:", err)
			break
		}

		// Process user input
		userInput := strings.TrimSpace(strings.ToUpper(string(message)))

		switch sessionState.Mode {
		case PromptMode:
			handlePromptInput(conn, sessionState, userInput)
		case PasswordMode:
			handlePasswordInput(conn, sessionState, userInput)
		case TicTacToeMode:
			handleTicTacToeInput(conn, sessionState, userInput)
		case GlobalThermonuclearWarMode:
			handleGlobalThermonuclearWarInput(conn, sessionState, userInput)
		}
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Print("Healthcheck Endpoint Hit")
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/wopr", handleNewConnection)
	r.HandleFunc("/wopr/{sessionID}", handleResumeConnection)
	r.HandleFunc("/healthcheck", healthcheck)

	fmt.Println("WEBSOCKET SERVER STARTED ON :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
