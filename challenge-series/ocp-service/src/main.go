package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// Globals
var endpoint = "/diagnostics"
var port = "8080"
var greetings = []string{".oO(Enforcement Droid Series 209 Support System)Oo.", "Interaction with this system limited to Omni Consumer Products employees and contractors only.", "Individuals suspected to have accessed this system without authorization are subject to prosecution or enrollment in OCP Indentured Employment program."}
var costFlag = 100
var costPerSecond = 1
var flag = "gc24{ff4a8d57-afe9-4cc2-a17f-769d56e3f7c4}"
var clients = make(map[*websocket.Conn]*sessionData)

// Types
type equation struct {
	Question        string
	RealAnswer      int
	Value           int
	PresentedAnswer int
}

type sessionData struct {
	Points          int
	CurrentEquation equation
	SessionID       string
	MessageIndex    int
	CycleCount      int
}

func (s *sessionData) addPoints(i int) {
	s.Points = s.Points + i
}

func (s *sessionData) subPoints(i int) {
	s.Points = s.Points - i
}

// generateSessionID creates random session identifiers for each client connection
func generateSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Disables CORS protection.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sendQuestion(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s = %s", clients[ws].CurrentEquation.Question, strconv.Itoa(clients[ws].CurrentEquation.PresentedAnswer))))
}
func sendInvalidInput(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("Error: Invalid input"))
}
func sendCorrect(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("Calibrating... Offset decreasing."))
}
func sendIncorrect(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("Calibrating... Offset increasing."))
}

// checkAnswer checks the client's current question against the answer they submitted.
func checkAnswer(ws *websocket.Conn, p string) {
	correctResponses := map[string]bool{"true": true, "correct": true, "yes": true, "y": true, "t": true}
	realAnswer := clients[ws].CurrentEquation.RealAnswer

	p = strings.ToLower(strings.TrimSpace(p))
	if correctResponses[p] {
		if clients[ws].CurrentEquation.RealAnswer == clients[ws].CurrentEquation.PresentedAnswer {
			sendCorrect(ws)
			clients[ws].addPoints(clients[ws].CurrentEquation.Value)
			log.Infof("Client confirmed correct answer! New point balance: %v", clients[ws].Points)
		} else {
			log.Infof("Incorrect affirmation received: %s", p)
			sendIncorrect(ws)
			clients[ws].subPoints(clients[ws].CurrentEquation.Value)
		}
	} else {
		guessedAnswer, err := strconv.Atoi(p)
		if err != nil {
			sendInvalidInput(ws)
			log.Errorf("Received error while converting input to int for comparison: %s", err)
			return
		}
		if guessedAnswer == realAnswer {
			clients[ws].addPoints(clients[ws].CurrentEquation.Value)
			log.Infof("Client provided correct answer! New point balance: %v", clients[ws].Points)
			sendCorrect(ws)
		} else {
			sendIncorrect(ws)
			clients[ws].subPoints(clients[ws].CurrentEquation.Value)
		}
	}
	if clients[ws].Points >= costFlag {
		ws.WriteMessage(websocket.TextMessage, []byte("Flag: "+flag))
		ws.WriteMessage(websocket.TextMessage, []byte("Goodbye!"))
		ws.Close()
	}
	clients[ws].CurrentEquation = generateEquation(clients[ws].Points / 10)
	sendQuestion(ws)
}

// handleConnection is the handler for websocket connections.
func handleConnection(w http.ResponseWriter, r *http.Request) {
	log.Info("Connection started...")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Error during connection upgrade: %s", err)
		return
	}
	sessionData := &sessionData{
		CurrentEquation: generateEquation(1),
		Points:          0,
		SessionID:       generateSessionID(),
	}

	clients[ws] = sessionData
	log.Info("Session Data established.")
	defer func() {
		delete(clients, ws)
		ws.Close()
	}()

	for _, greeting := range greetings {
		err = ws.WriteMessage(websocket.TextMessage, []byte(greeting))
		if err != nil {
			log.Errorf("Error sending greeting: %s", err)
		}
	}
	sendQuestion(ws)
	log.Info("Greeting Sent, listening for requests.")

	for {
		if clients[ws].Points >= costFlag {
			ws.WriteMessage(websocket.TextMessage, []byte("Calibration completed. Generating repair receipt."))
			ws.WriteMessage(websocket.TextMessage, []byte(flag))
			ws.WriteMessage(websocket.TextMessage, []byte("Goodbye!"))
			ws.Close()
		}
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Errorf("WebSocket closed: %s", err)
			} else {
				log.Errorf("Error reading from client: %s", err)
			}
			return
		}

		if messageType == websocket.TextMessage {
			log.Infof("Received message from client: %s", p)
			checkAnswer(ws, string(p))
		}
	}

}

// pointDecay is run in a goroutine and handles the process of removing points from client connections over time.
// Purpose of this behavior is to discourage people just using a websocket terminal and slowly solving math by hand.
// Piggybacking on this to shoehorn in background noise emission to clients.
func pointDecay() {
	log.Infof("Pointdecay process started.")
	for {
		for ws, v := range clients {
			v.subPoints(costPerSecond)

			// Check if it's time to send a message
			if v.CycleCount >= 120 && (v.CycleCount == 0 || v.CycleCount >= 20) {
				currentMessage := getNextMessage(v) // Get the next message directly
				if err := ws.WriteMessage(websocket.TextMessage, []byte(currentMessage)); err != nil {
					log.Errorf("Error sending message: %v", err)
					continue
				}
				v.CycleCount = 1 // Reset cycle count for this client
			} else {
				v.CycleCount++ // Increment cycle count if not sending a message
			}
		}
		time.Sleep(1 * time.Second)
		log.Infof("Tick")
	}
}

func getNextMessage(session *sessionData) string {
	messages := []string{
		"AUDIO LOG 2B22: I think you'd better do what he says, Mr. Kinney.",
		"RADIO INTERCEPT 6F17: All units, 211 in progress.",
		"INCIDENT LOG 036B: Suspect non-compliance. Physical force authorized.",
		"AUDIO LOG 442F: Oh my god!",
		"RADIO INTERCEPT 8FE9: Where's that backup?!.",
		"AUDIO LOG 0CD7: Somebody call a paramedic!",
		"RADIO INTERCEPT 39B0: No backup available.",
		"INCIDENT LOG E502: Unit motion fault detected.",
	}

	if session.MessageIndex < len(messages) {
		message := messages[session.MessageIndex]
		session.MessageIndex += 1
		return message
	}
	return "RADIO INTERCEPT: *static*"
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Infof("Healthcheck Endpoint Hit")
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	http.HandleFunc(endpoint, handleConnection)
	http.HandleFunc("/healthcheck", healthcheck)

	log.Infof("WebSocket server started at ws://localhost:%s%s", port, endpoint)
	go pointDecay()
	http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		nil)
}
