package dto

type PlayerGroup struct {
	PlayerId   int     `json:"playerid"`
	GroupId    int     `json:"groupid"`
	Editor     []uint8 `json:"editor"`
	PlayerName string  `json:"playername"`
	GroupName  string  `json:"groupname"`
}
