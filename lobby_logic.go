package main

import (
	"sync"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var (
	createDeleteMutex = &sync.Mutex{}
	lobbies           = []*Lobby{}
)

// Player represents a participant in a lobby.
type Player struct {
	// Name is the players displayed name
	Name string
	// UserSession uniquely identifies the player.
	UserSession string
	// Score is the points that the player got in the current lobby.
	Score int
	Rank  int
	ws    *websocket.Conn
}

// Lobby represents a game session.
type Lobby struct {
	// ID uniquely identified the Lobby.
	ID string

	// Password defines the password that participants need to enter before
	// being allowed to join the Lobby.
	Password string
	// DrawingTime is the amount of seconds that each player has available to
	// finish their drawing.
	DrawingTime int
	// Rounds defines how many iterations a lobby does before the game ends.
	// One iteration means every participant does one drawing.
	Rounds int
	// MaxPlayers defines the maximum amount of players in a single lobby.
	MaxPlayers int
	// CustomWords are additional words that will be used in addition to the
	// predefined words.
	CustomWords []string

	// Players references all participants of the Lobby.
	Players []*Player

	// Drawer references the Player that is currently drawing.
	Drawer *Player
	// CurrentWord represents the word that was last selected. If no word has
	// been selected yet or the round is already over, this should be empty.
	CurrentWord string
	// WordHints for the current word.
	WordHints []*WordHint
	// Round is the round that the Lobby is currently in. This is a number
	// between 0 and Rounds. 0 indicates that it hasn't started yet.
	Round int
}

// GetPlayer searches for a player, identifying them by usersssion.
func (lobby *Lobby) GetPlayer(userSession string) *Player {
	for _, player := range lobby.Players {
		if player.UserSession == userSession {
			return player
		}
	}

	return nil
}

// GetLobby returns a Lobby that has a matching ID or no Lobby if none could
// be found.
func GetLobby(id string) *Lobby {
	for _, l := range lobbies {
		if l.ID == id {
			return l
		}
	}

	return nil
}

func createLobby(
	password string,
	drawingTime int,
	rounds int,
	maxPlayers int,
	customWords []string) *Lobby {

	createDeleteMutex.Lock()

	lobby := &Lobby{
		ID:          uuid.NewV4().String(),
		Password:    password,
		DrawingTime: drawingTime,
		Rounds:      rounds,
		MaxPlayers:  maxPlayers,
		CustomWords: customWords,
	}

	lobbies = append(lobbies, lobby)

	createDeleteMutex.Unlock()

	return lobby
}

func createPlayer(name string) *Player {
	return &Player{
		Name:        name,
		UserSession: uuid.NewV4().String(),
		Score:       0,
	}
}