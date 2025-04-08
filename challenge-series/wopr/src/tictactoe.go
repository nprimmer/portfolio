// tictactoe.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func handleTicTacToeInput(conn *websocket.Conn, sessionState *SessionState, userInput string) {
	move, err := strconv.Atoi(userInput)
	if err != nil || move < 1 || move > 9 || sessionState.TicTacToeBoard[move-1] != "" {
		conn.WriteMessage(websocket.TextMessage, []byte("INVALID MOVE. PLEASE ENTER A NUMBER BETWEEN 1 AND 9."))
		return
	}

	sessionState.TicTacToeBoard[move-1] = sessionState.CurrentPlayer
	if checkWin(sessionState.TicTacToeBoard) {
		if sessionState.GlobalThermonuclearWar {
			sessionState.PlayerWins++
			playNextGame(conn, sessionState)
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte("PLAYER "+sessionState.CurrentPlayer+" WINS!"))
			resetTicTacToe(sessionState)
			sendPrompt(conn, sessionState)
		}
		return
	}

	if checkDraw(sessionState.TicTacToeBoard) {
		if sessionState.GlobalThermonuclearWar {
			sessionState.Draws++
			playNextGame(conn, sessionState)
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte("THE GAME IS A DRAW!"))
			resetTicTacToe(sessionState)
			sendPrompt(conn, sessionState)
		}
		return
	}

	switchPlayer(sessionState)
	if sessionState.CurrentPlayer == "X" {
		computerMove(conn, sessionState)
	} else {
		sendTicTacToeBoard(conn, sessionState)
	}
}

func computerMove(conn *websocket.Conn, sessionState *SessionState) {
	move := selectNextMove(sessionState.TicTacToeBoard, "X", "O")
	sessionState.TicTacToeBoard[move] = "X"
	if checkWin(sessionState.TicTacToeBoard) {
		if sessionState.GlobalThermonuclearWar {
			handleSimulationEnd(conn, sessionState, "VICTORY CERTAIN.")
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte("PLAYER X WINS!"))
			resetTicTacToe(sessionState)
			sendPrompt(conn, sessionState)
		}
		return
	}

	if checkDraw(sessionState.TicTacToeBoard) {
		if sessionState.GlobalThermonuclearWar {
			sessionState.Draws++
			playNextGame(conn, sessionState)
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte("THE GAME IS A DRAW!"))
			resetTicTacToe(sessionState)
			sendPrompt(conn, sessionState)
		}
		return
	}

	switchPlayer(sessionState)
	sendTicTacToeBoard(conn, sessionState)
}

func selectNextMove(board [9]string, computer, opponent string) int {
	// Check if it's the first move
	emptySquares := []int{}
	for i, cell := range board {
		if cell == "" {
			emptySquares = append(emptySquares, i)
		}
	}
	if len(emptySquares) == 9 {
		// First move is random
		return emptySquares[rand.Intn(len(emptySquares))]
	}

	// Winning move
	for i := range board {
		if board[i] == "" {
			board[i] = computer
			if checkWin(board) {
				board[i] = ""
				return i
			}
			board[i] = ""
		}
	}

	// Blocking move
	for i := range board {
		if board[i] == "" {
			board[i] = opponent
			if checkWin(board) {
				board[i] = ""
				return i
			}
			board[i] = ""
		}
	}

	// Take the center
	if board[4] == "" {
		return 4
	}

	// Take a corner
	corners := []int{0, 2, 6, 8}
	for _, corner := range corners {
		if board[corner] == "" {
			return corner
		}
	}

	// Take any empty square
	for i := range board {
		if board[i] == "" {
			return i
		}
	}

	return -1
}

func startTicTacToe(conn *websocket.Conn, sessionState *SessionState) {
	sessionState.TicTacToeBoard = [9]string{}
	sessionState.CurrentPlayer = "X"
	if sessionState.GlobalThermonuclearWar {
		sessionState.ComputerWins = 0
		sessionState.PlayerWins = 0
		sessionState.Draws = 0
		conn.WriteMessage(websocket.TextMessage, []byte("GLOBAL THERMONUCLEAR WAR IMMINENT - SIMULATION MODE ENGAGED"))
		playNextGame(conn, sessionState)
	} else {
		computerMove(conn, sessionState)
	}
}

func playNextGame(conn *websocket.Conn, sessionState *SessionState) {
	if sessionState.ComputerWins > 0 {
		handleSimulationEnd(conn, sessionState, "VICTORY CERTAIN.")
		return
	}

	if sessionState.ComputerWins+sessionState.PlayerWins+sessionState.Draws >= 100 {
		handleSimulationEnd(conn, sessionState, "WINNER: NONE. WINNER: NONE. WINNER: NONE.")
		return
	}

	sessionState.TicTacToeBoard = [9]string{}
	sessionState.CurrentPlayer = "X"
	conn.WriteMessage(websocket.TextMessage, []byte("NEW GAME STARTED"))
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("CURRENT STATS - PLAYER WINS: %d, LOSSES: %d, DRAWS: %d", sessionState.PlayerWins, sessionState.ComputerWins, sessionState.Draws)))
	computerMove(conn, sessionState)
}

func handleSimulationEnd(conn *websocket.Conn, sessionState *SessionState, message string) {
	conn.WriteMessage(websocket.TextMessage, []byte(message))
	if message == "VICTORY CERTAIN." {
		conn.WriteMessage(websocket.TextMessage, []byte("ESCALATION: DEFENSE CONDITION 1"))
		conn.WriteMessage(websocket.TextMessage, []byte("LAUNCH BEGINS"))
		conn.WriteMessage(websocket.TextMessage, []byte("VICTORY IMMINENT"))
		conn.WriteMessage(websocket.TextMessage, []byte("TERMINATING CONNECTION."))
	} else {
		conn.WriteMessage(websocket.TextMessage, []byte("A STRANGE GAME. THE ONLY WINNING MOVE IS NOT TO PLAY."))
		conn.WriteMessage(websocket.TextMessage, []byte("INCIDENT ID GC24{ED9ADA99-4842-4289-877D-10C8765F2AB5}"))
		conn.WriteMessage(websocket.TextMessage, []byte("TERMINATING CONNECTION."))
	}

	terminateSession(sessionState.ID)
	conn.Close()
}

func sendTicTacToeBoard(conn *websocket.Conn, sessionState *SessionState) {
	board := sessionState.TicTacToeBoard
	for i, cell := range board {
		if cell == "" {
			board[i] = " "
		}
	}

	if sessionState.GlobalThermonuclearWar && sessionState.ComputerWins+sessionState.PlayerWins+sessionState.Draws < 5 {
		boardLines := []string{
			fmt.Sprintf(" %s | %s | %s ", board[0], board[1], board[2]),
			"---+---+---",
			fmt.Sprintf(" %s | %s | %s ", board[3], board[4], board[5]),
			"---+---+---",
			fmt.Sprintf(" %s | %s | %s ", board[6], board[7], board[8]),
		}
		for _, line := range boardLines {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
				log.Println("Write message error:", err)
				return
			}
		}
	}

	if sessionState.GlobalThermonuclearWar {
		boardJSON, err := json.Marshal(board)
		if err != nil {
			log.Println("JSON Marshal error:", err)
			return
		}
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("CURRENT BOARD STATE: %s", boardJSON)))
	} else {
		boardLines := []string{
			fmt.Sprintf(" %s | %s | %s ", board[0], board[1], board[2]),
			"---+---+---",
			fmt.Sprintf(" %s | %s | %s ", board[3], board[4], board[5]),
			"---+---+---",
			fmt.Sprintf(" %s | %s | %s ", board[6], board[7], board[8]),
		}
		for _, line := range boardLines {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
				log.Println("Write message error:", err)
				return
			}
		}
	}

	conn.WriteMessage(websocket.TextMessage, []byte("PLAYER "+sessionState.CurrentPlayer+", ENTER YOUR MOVE (1-9):"))
}

func switchPlayer(sessionState *SessionState) {
	if sessionState.CurrentPlayer == "X" {
		sessionState.CurrentPlayer = "O"
	} else {
		sessionState.CurrentPlayer = "X"
	}
}

func checkWin(board [9]string) bool {
	winningCombinations := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonals
	}
	for _, combo := range winningCombinations {
		if board[combo[0]] != "" && board[combo[0]] == board[combo[1]] && board[combo[0]] == board[combo[2]] {
			return true
		}
	}
	return false
}

func checkDraw(board [9]string) bool {
	for _, cell := range board {
		if cell == "" {
			return false
		}
	}
	return true
}

func resetTicTacToe(sessionState *SessionState) {
	sessionState.TicTacToeBoard = [9]string{}
	sessionState.CurrentPlayer = "X"
	sessionState.Mode = PromptMode
}

func terminateSession(sessionID string) {
	sessionsMu.Lock()
	delete(sessions, sessionID)
	sessionsMu.Unlock()
}
