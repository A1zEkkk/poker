package handler

import (
	"poker/room/roommanager"
)

type RoomInfoResponse struct {
	ID         string   `json:"id"`
	HostID     string   `json:"hostId"`
	State      string   `json:"state"`
	MaxPlayers int      `json:"maxPlayers"`
	MinBank    float64  `json:"minBank"`
	PlayerIDs  []string `json:"players"`
}

func RoomToInfoResponse(r *roommanager.Room) RoomInfoResponse {
	playerIDs := make([]string, 0, len(r.Players))
	for id := range r.Players {
		playerIDs = append(playerIDs, id)
	}

	return RoomInfoResponse{
		ID:         r.ID,
		HostID:     r.HostID,
		State:      string(r.State),
		MaxPlayers: r.MaxPlayers,
		MinBank:    r.MinBank,
		PlayerIDs:  playerIDs,
	}
}

type JoinRoomRequest struct {
	RoomID string  `json:"room_id"`
	Bank   float64 `json:"bank"`
}
