package dto

/*
	/player/getplayer/
	Get a player's information.
*/
type GetPlayerRequest struct {
	PlayerId int `json:"playerid"`
}

/*
	/player/create/
	Creates a new player.
*/
type CreatePlayerRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

/*
	/player/getgroups/
	Gets all PlayerGroups that player is in.
*/
type GetAllGroups struct {
	PlayerId string `json:"playerid"`
}

/*
	/player/changename/
	Changes a player's name. Must be an editor user? Or account linked to player must match.
*/
type ChangePlayerNameRequest struct {
	Playerid  int    `json:"playerid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

/*
	/player/updatelifetimetotals/
	Update a player's lifetime totals.
*/
type UpdatePlayerLifetimeTotalsRequest struct {
	PlayerId      int `json:"playerid"`
	Kills         int `json:"kills"`
	Assists       int `json:"assists"`
	Aces          int `json:"aces"`
	Digs          int `json:"digs"`
	Blocks        int `json:"blocks"`
	AtkErrors     int `json:"atkerrors"`
	ServiceErrors int `json:"serviceerrors"`
	AssistErrors  int `json:"assisterrors"`
	BlockErrors   int `json:"blockerrors"`
}
