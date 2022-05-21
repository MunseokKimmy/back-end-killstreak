package dto

import "time"

type Group struct {
	GroupId           int
	Name              string
	DateCreated       time.Time
	GameLastCompleted time.Time
}
