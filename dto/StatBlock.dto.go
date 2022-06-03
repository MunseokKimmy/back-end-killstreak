package dto

type StatBlock struct {
	StatId        int
	GameId        int
	PlayerId      int
	Kills         int
	AtkErrors     int
	ServiceAces   int
	ServiceErrors int
	Assists       int
	AssistErrors  int
	Digs          int
	Blocks        int
	BlockErrors   int
	OnTeamOne     []uint8
	PlayerName    string
	GameName      string
}