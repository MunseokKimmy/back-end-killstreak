package dto

type Player struct {
	PlayerId      int
	FirstName     string
	LastName      string
	Kills         int
	Assists       int
	Aces          int
	Digs          int
	Blocks        int
	AtkErrors     int
	ServiceErrors int
	AssistErrors  int
	BlockErrors   int
}

type PlayerShort struct {
	PlayerId int
	Name     string
}
