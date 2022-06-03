package dto

import "time"

type Game struct {
	GameId         int
	GroupId        int
	Name           string
	Date           time.Time
	TeamOneName    string
	TeamTwoName    string
	TeamOneScore   int
	TeamTwoScore   int
	TeamOneServing []uint8
	LastUpdated    time.Time
	Completed      []uint8
}

type GameShort struct {
	GameId   int
	GameName string
}
