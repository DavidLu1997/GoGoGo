// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

type Move struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
	Id       string `json:"id"`
	X        uint   `json:"move_x"`
	Y        uint   `json:"move_y"`
	CreateAt int64  `json:"create_at"`
}

func (m *Move) ToJson() string {
	s, err := json.Marshal(m)
	if err != nil {
		return ""
	} else {
		return string(s)
	}
}

func (m *Move) PreSave() {
	m.CreateAt = GetMillis()

	if m.Id == "" {
		m.Id = NewId()
	}
}

func (m *Move) IsValid(game *Game) *Error {
	currentPiece, err := game.GetBoardPiece(m.X, m.Y)

	if err != nil {
		return err
	} else if currentPiece != 0 {
		return NewLocError("Move.IsValid", "Spot is occupied", nil, "")
	}
	return nil

}

func MoveFromJson(data io.Reader) *Move {
	decoder := json.NewDecoder(data)
	var m Move
	err := decoder.Decode(&m)
	if err == nil {
		return &m
	}
	return nil
}

func MovesToJson(m []*Move) string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func MovesFromJson(data io.Reader) []*Move {
	decoder := json.NewDecoder(data)
	var o []*Move
	err := decoder.Decode(&o)
	if err == nil {
		return o
	} else {
		return nil
	}
}
