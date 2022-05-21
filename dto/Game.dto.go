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
	TeamOneServing bool
	LastUpdated    time.Time
}
