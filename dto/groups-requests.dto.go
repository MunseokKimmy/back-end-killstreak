package dto

import "time"

/*
	Either use request or query params.
	/groups/getgroup/
	Grabs group data.
*/
type GetGroupRequest struct {
	Id int `json:"groupid"`
}

/*
	/groups/getplayers/
	Gets all groups a player is in.
*/
type GetGroupsOfPlayer struct {
	PlayerId int `json:"playerid"`
}

/*
	/groups/creategroup/
	Creates a new group and makes the player creating the group an editor user.
*/
type CreateGroupRequest struct {
	Name       string `json:"name"`
	PlayerId   int    `json:"playerid"`
	PlayerName string `json:"playername"`
}

/*
	/groups/addplayer/
	Adds an existing player to the group. Must be an editor user.
*/
type AddPlayerRequest struct {
	GroupId    int    `json:"groupid"`
	PlayerId   int    `json:"playerid"`
	EditorId   int    `json:"editorid"`
	GroupName  string `json:"groupname"`
	PlayerName string `json:"playername"`
}

/*
	/groups/updatelastcompletedgame/
	Updates a group's "LastCompletedGame" date to the new date.
*/
type UpdateLastCompletedGameRequest struct {
	GroupId int       `json:"groupid"`
	NewDate time.Time `json:"newdate"`
}

/*
	/groups/giveplayereditor/
	Give another player editor permissions in the group. Must be an editor user.
*/
type GivePlayerEditorRequest struct {
	GroupId               int `json:"groupid"`
	NewEditorPlayerId     int `json:"newEditorId"`
	CurrentEditorPlayerId int `json:"currentEditorId"`
}

/*
	/groups/changename/
	Change name of the group. Must be an editor user.
*/
type ChangeNameRequest struct {
	GroupId  int    `json:"groupid"`
	EditorId int    `json:"editorid"`
	Name     string `json:"name"`
}

/*
	/groups/deletegroup/
	Delete group. Must be an editor user.
*/
type DeleteGroupRequest struct {
	GroupId  int `json:"groupid"`
	EditorId int `json:"editorid"`
}

/*
	/groups/removeplayer/
	Remove player from the group. Must be an editor user.
*/
type RemovePlayerRequest struct {
	GroupId             int `json:"groupid"`
	PlayerToBeRemovedId int `json:"removedplayerid"`
	EditorId            int `json:"editorid"`
}
