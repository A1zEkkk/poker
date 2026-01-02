package room

import (
	"encoding/json"
	"net/http"
)

type RoomInfo struct { //Get data
	ID             string   `json:"id"`
	HostId         string   `json:"host_id"`
	MinBank        float64  `json:"min_bank"`
	MaxPlayers     int      `json:"max_players"`
	CurrentPlayers int      `json:"current_players"`
	PlayersIDs     []string `json:"players"`
}

type CreateRoomRequest struct { //Post data
	MaxPlayers int     `json:"max_players"`
	HostId     string  `json:"host_id"`
	MinBank    float64 `json:"min_bank"`
}

type JoinRoomRequest struct {
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
}

type LeaveRoomRequest struct {
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
}

// У нас должен быть jwt токен + айдишник комнаты

// После нажатия на кнопку комнаты у нас возвращается всся информация о комнате, в которую мы в теории можем зайти
func (rm *RoomManager) GetRoomHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что мы вызвали нужный метод ибо к 1 юр можем привязать несколько методов
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получаем id комнаты информацию, о которой будем передавать
	roomID := r.URL.Query().Get("id")
	if roomID == "" {
		http.Error(w, "missing room id", http.StatusBadRequest)
		return
	}

	//Получаем комнату из множества комнат
	room, err := rm.GetRoom(roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Собираем список игроков
	room.mu.Lock()
	playerIDs := make([]string, len(room.Game.Players))
	for i, p := range room.Game.Players {
		playerIDs[i] = p.Id
	}
	room.mu.Unlock()

	//Генерируем Json, который передадим клиенту
	response := RoomInfo{
		ID:             room.ID,
		HostId:         room.HostId,
		MinBank:        room.MinBank,
		MaxPlayers:     room.MaxPlayers,
		CurrentPlayers: room.CurrentPlayers,
		PlayersIDs:     playerIDs,
	}

	//Как я понял тут нету такого понятия, возвращаемое значени ф-ции и го передает через этот интерфес ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//В теории нам будет достаточно произвести провеку jwt токена. Будет реализован позже

func (rm *RoomManager) GetListIdRoomHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что мы вызвали нужный метод ибо к 1 юрл можем привязать несколько методов
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rooms := rm.ListRooms()
	romsId := make([]string, len(rooms))
	for i := 0; i < len(rooms); i++ {
		romsId[i] = rooms[i].ID
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(romsId)
}

func (rm *RoomManager) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateRoomRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	//Валидируем данные
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.MaxPlayers <= 0 {
		http.Error(w, "max_players must be > 0", http.StatusBadRequest)
		return
	}

	if req.HostId == "" {
		http.Error(w, "host_id required", http.StatusBadRequest)
		return
	}

	// Создаем комнату
	room, err := rm.CreateRoom(req.MaxPlayers, req.HostId, req.MinBank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := RoomInfo{
		ID:             room.ID,
		HostId:         room.HostId,
		MinBank:        room.MinBank,
		MaxPlayers:     room.MaxPlayers,
		CurrentPlayers: room.CurrentPlayers,
		PlayersIDs:     room.PlayersIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (rm *RoomManager) JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req JoinRoomRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	//Валидируем данные
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.RoomID == "" {
		http.Error(w, "RoomID required", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "user_id required", http.StatusBadRequest)
		return
	}

	err = rm.JoinRoom(req.RoomID, req.UserID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("joined"))
}

func (rm *RoomManager) LeaveRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LeaveRoomRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	//Валидируем данные
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.RoomID == "" {
		http.Error(w, "RoomID required", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "user_id required", http.StatusBadRequest)
		return
	}

	err = rm.LeaveRoom(req.RoomID, req.UserID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("leaved"))
}
