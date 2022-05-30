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

/*
	/game/creategame/
	Creates a new game, giving it a name and both teams names. Needs a group?
	Requires editorid.
*/
type CreateGameRequest struct {
	GroupId     int    `json:"groupid,omitempty"`
	EditorId    int    `json:"editorid"`
	Name        string `json:"name"`
	TeamOneName string `json:"teamonename"`
	TeamTwoName string `json:"teamtwoname"`
}

/*
	/game/changename/
	Changes the name of a game. Requires editorId.
*/
type GameChangeNameRequest struct {
	GameId   int    `json:"gameid"`
	NewName  string `json:"newname"`
	EditorId int    `json:"editorid"`
}

/*
	/game/teamonescore/
	Increments team one's score. Requires editorId.
*/
type TeamOneScoreRequest struct {
	GameId   int `json:"gameid"`
	EditorId int `json:"editorid"`
}

/*
	/game/teamtwoscore/
	Increments team two's score. Requires editorId.
*/
type TeamTwoScoreRequest struct {
	GameId   int `json:"gameid"`
	EditorId int `json:"editorid"`
}

/*
	/game/completegame/
	Marks game as complete. Requires editorId.
*/
type GameCompletedRequest struct {
	GameId   int `json:"gameid"`
	EditorId int `json:"editorid"`
}

/*
	/game/switchserver/
	Chooses serving team. Requires editorId.
*/
type GameSwitchServer struct {
	GameId         int  `json:"gameid"`
	EditorId       int  `json:"editorid"`
	TeamOneServing bool `json:"teamoneserving"`
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
}
