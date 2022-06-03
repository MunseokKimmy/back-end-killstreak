package dto

import "time"

/*

 */
type GetGameRequest struct {
	GameId int `json:"gameid"`
}

type GetAllGamesOfPlayer struct {
	PlayerId int `json:"playerid"`
}

type GetAllPlayersInGame struct {
	GameId int `json:"gameid"`
}

/*
	/game/creategame/
	Creates a new game, giving it a name and both teams names. Needs a group?
	Requires editorid.
*/
type CreateGameRequest struct {
	GroupId     int    `json:"groupid,omitempty"`
	EditorId    int    `json:"editorid"`
	EditorName  string `json:"editorname"`
	Name        string `json:"name"`
	TeamOneName string `json:"teamonename"`
	TeamTwoName string `json:"teamtwoname"`
	Players     []int  `json:"playerids,omitempty"`
	EditorTeam  int    `json:"editorteam"` //2 - not in game, 1 - team 1, 0 - team 2
}

const (
	EditorTeamOne    int = 0
	EditorTeamTwo    int = 1
	EditorNotPlaying int = 2
)

/*
	/game/changename/
	Changes the name of a game. Requires editorId.
*/
type GameChangeNameRequest struct {
	GameId   int    `json:"gameid"`
	NewName  string `json:"newname"`
	EditorId int    `json:"editorid"`
	GroupId  int    `json:"groupid"`
}

/*
	/game/addplayer/
	Adds a player to a game. Requires editorId.
*/
type AddPlayerToGameRequest struct {
	PlayerId   int    `json:"playerid"`
	PlayerName string `json:"playername"`
	EditorId   int    `json:"editorid"`
	GroupId    int    `json:"groupid"`
	GameId     int    `json:"gameid"`
	GameName   string `json:"gamename"`
	OnTeamOne  int    `json:"onteamone"`
}

/*
	/game/removeplayer/
	Removes a player from a game. Requires editorId.
*/
type RemovePlayerFromGameRequest struct {
	PlayerId int `json:"playerid"`
	EditorId int `json:"editorid"`
	GameId   int `json:"gameid"`
	GroupId  int `json:"groupid"`
}

/*
	/game/teamonescore/
	Increments team one's score. Requires editorId.
*/
type TeamOneScoreRequest struct {
	GameId   int `json:"gameid"`
	EditorId int `json:"editorid"`
	GroupId  int `json:"groupid"`
}

/*
	/game/teamtwoscore/
	Increments team two's score. Requires editorId.
*/
type TeamTwoScoreRequest struct {
	GameId   int `json:"gameid"`
	EditorId int `json:"editorid"`
	GroupId  int `json:"groupid"`
}

/*
	/game/completegame/
	Marks game as complete. Requires editorId.
*/
type GameCompletedRequest struct {
	GameId   int `json:"gameid"`
	EditorId int `json:"editorid"`
	GroupId  int `json:"groupid"`
}

/*
	/game/uncompletegame/
	Marks game as complete. Requires editorId.
*/
type GameInCompletedRequest struct {
	GameId   int `json:"gameid"`
	EditorId int `json:"editorid"`
	GroupId  int `json:"groupid"`
}

/*
	/game/switchserver/
	Chooses serving team. Requires editorId.
*/
type GameSwitchServer struct {
	GameId         int  `json:"gameid"`
	EditorId       int  `json:"editorid"`
	TeamOneServing bool `json:"teamoneserving"`
	GroupId        int  `json:"groupid"`
}

/*
	/game/update/
	Updates the entire Game row. Requires editor user.
*/
type UpdateGameRequest struct {
	GameId         int       `json:"gameid"`
	EditorId       int       `json:"editorid"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
	TeamOneName    string    `json:"teamonename"`
	TeamTwoName    string    `json:"teamtwoname"`
	TeamOneScore   int       `json:"teamonescore"`
	TeamTwoScore   int       `json:"teamtwoscore"`
	TeamOneServing bool      `json:"teamoneserving"`
	GroupId        int       `json:"groupid"`
}
