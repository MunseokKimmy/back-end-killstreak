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
	Team          int
	PlayerName    string
	GameName      string
}
