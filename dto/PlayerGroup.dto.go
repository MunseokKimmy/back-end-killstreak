package dto

type PlayerGroup struct {
	PlayerId   int    `json:"playerid"`
	GroupId    int    `json:"groupid"`
	Editor     any    `json:"editor"`
	PlayerName string `json:"playername"`
	GroupName  string `json:"groupname"`
}
