package dto

/*
	/stats/getgamestats/
	Gets all stats for a game.
*/
type GetGameStatsRequest struct {
	GameId int `json:"gameid"`
}

/*
	/stats/getplayerstatsingame/
	Gets all stats for a player in a game.
*/
type GetPlayersStatsInGameRequest struct {
	GameId   int `json:"gameid"`
	PlayerId int `json:"playerid"`
}

/*
	/stats/getplayerstats/
	Gets all stats for a player.
*/
type GetPlayersStatsRequest struct {
	PlayerId int `json:"playerid"`
}

/*
	/stats/create/
	Creates a new set of stats. Used when a game is created with a player in it.
	Requires editor user.
*/
type CreateStatsRequest struct {
	EditorId   int    `json:"editorid"`
	GameId     int    `json:"gameid"`
	PlayerId   int    `json:"playerid"`
	GroupId    int    `json:"groupid"`
	GameName   string `json:"gamename"`
	PlayerName string `json:"playername"`
}

/*
	/stats/update/
	Updates block of stats for a player. Used during the game for any stat recorded.
	Requires editor user.
*/
type UpdateStatsRequest struct {
	EditorId      int `json:"editorid"`
	GameId        int `json:"gameid"`
	PlayerId      int `json:"playerid"`
	GroupId       int `json:"groupid"`
	Kills         int `json:"kills"`
	AtkErrors     int `json:"atkerrors"`
	ServiceAces   int `json:"serviceaces"`
	ServiceErrors int `json:"serviceerrors"`
	Assists       int `json:"assists"`
	AssistErrors  int `json:"assisterrors"`
	Digs          int `json:"digs"`
	Blocks        int `json:"blocks"`
	BlockErrors   int `json:"blockerrors"`
}

/*
	/stats/highlights/
	Gets all group stats and adds them together.
*/
type HighlightsRequest struct {
	GroupId int    `json:"groupid"`
	Date    string `json:"date"`
}

/*
	/stats/allgroupstats/
	Gets Group Stats On a certain day.
*/
type AllGroupStatsOnDayRequest struct {
	GroupId int    `json:"groupid"`
	Date    string `json:"date"`
}
