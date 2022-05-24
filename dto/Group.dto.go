package dto

import "time"

type Group struct {
	GroupId           int
	Name              string
	DateCreated       time.Time
	GameLastCompleted time.Time
}

type GroupShort struct {
	GroupId int    `json:"groupid"`
	Name    string `json:"groupname"`
}
