package roommanager

import (
	"errors"
	"poker/game"
	"sync"
)

type RoomState string

const (
	RoomWaiting RoomState = "waiting"
	RoomPlaying RoomState = "playing"
)

type RoomPlayer struct { // Реализовать функцию, которая будет получать данные о пользователе(репозиторий) +
	ID    string
	Bank  float64
	Ready bool
	// Conn / Session добавим позже
}

type Room struct {
	ID         string
	HostID     string
	MaxPlayers int
	MinBank    float64

	State   RoomState
	Players map[string]*RoomPlayer
	Game    *game.Game

	Mu sync.Mutex
}

func NewRoom(id, hostID string, maxPlayers int, minBank float64) *Room {
	return &Room{
		ID:         id,
		HostID:     hostID,
		MaxPlayers: maxPlayers,
		MinBank:    minBank,
		State:      RoomWaiting,
		Players:    make(map[string]*RoomPlayer),
	}
}

func (r *Room) CanJoin(playerID string, bank float64) error {
	if r.State != RoomWaiting {
		return errors.New("game already started")
	}
	if len(r.Players) >= r.MaxPlayers {
		return errors.New("room is full")
	}
	if bank < r.MinBank {
		return errors.New("bank too small")
	}
	if _, exists := r.Players[playerID]; exists {
		return errors.New("already in room")
	}
	return nil
}

func (r *Room) Join(playerID string, bank float64) error {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	if err := r.CanJoin(playerID, bank); err != nil {
		return err
	}

	r.Players[playerID] = &RoomPlayer{
		ID:   playerID,
		Bank: bank,
	}
	return nil
}

func (r *Room) Leave(playerID string) {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	delete(r.Players, playerID)
}

func (r *Room) ToggleReady(playerID string) (bool, error) {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	if r.State != RoomWaiting {
		return false, errors.New("cannot change ready state during game")
	}

	p, ok := r.Players[playerID]
	if !ok {
		return false, errors.New("player not in room")
	}

	p.Ready = !p.Ready
	return p.Ready, nil
}

func (r *Room) CanStart() bool {
	if r.State != RoomWaiting {
		return false
	}

	if len(r.Players) < 2 {
		return false
	}

	readyCount := 0
	for _, p := range r.Players {
		if p.Ready {
			readyCount++
		}
	}

	return readyCount == len(r.Players)
}

func (r *Room) JoinRoom(playerID string, bank float64) error {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	if r.State != RoomWaiting {
		return errors.New("game already started")
	}

	if len(r.Players) >= r.MaxPlayers {
		return errors.New("room is full")
	}

	if bank < r.MinBank {
		return errors.New("bank too small")
	}

	if _, exists := r.Players[playerID]; exists {
		return errors.New("already in room")
	}

	r.Players[playerID] = &RoomPlayer{
		ID:   playerID,
		Bank: bank,
	}

	return nil
}
