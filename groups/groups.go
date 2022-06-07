package groups

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"killstreak/dto"
	"killstreak/utils"
	"net/http"
	"strconv"
)

// /groups/
func GetAllGroups(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.Group, error) {
	rows, err := db.Query("SELECT * FROM groups")
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()

	groups := make([]*dto.Group, 0)
	for rows.Next() {
		group := new(dto.Group)
		err := rows.Scan(&group.GroupId, &group.Name, &group.DateCreated, &group.GameLastCompleted)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		groups = append(groups, group)
	}
	err = rows.Err()
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return groups, nil
}

// /groups/get?groupid=XXXX
func GetGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) (*dto.Group, error) {
	var err error
	id := r.URL.Query().Get("groupid")
	group := new(dto.Group)
	row := db.QueryRow("SELECT * FROM groups WHERE groupid = ?;", id)
	err = row.Scan(&group.GroupId, &group.Name, &group.DateCreated, &group.GameLastCompleted)
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return group, nil
}

// /groups/player
func GetAllGroupsOfPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.GroupShort, error) {
	var request dto.GetGroupsOfPlayer
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}

	rows, err := db.Query("SELECT groupid, groupname FROM playergroup where playerid = ?;", strconv.Itoa(request.PlayerId))
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()

	groups := make([]*dto.GroupShort, 0)
	for rows.Next() {
		group := new(dto.GroupShort)
		err := rows.Scan(&group.GroupId, &group.Name)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		groups = append(groups, group)
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return groups, nil
}

// /groups/create
func CreateGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.CreateGroupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	result, err := db.Exec("INSERT INTO groups (name) VALUES (?);", request.Name)
	if utils.Error500Check(err, w) {
		return err
	}
	id, err := result.LastInsertId()
	if utils.Error500Check(err, w) {
		return err
	}
	_, err = db.Exec("INSERT INTO playergroup (groupid, playerid, groupname, playername, editor) VALUES (?, ?, ?, ?, ?);", strconv.FormatInt(id, 10), strconv.Itoa(request.PlayerId), request.Name, request.PlayerName, 1)
	if utils.Error500Check(err, w) {
		return err
	}
	return nil
}

// /groups/addplayer
func AddPlayerToGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.AddPlayerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil
	}
	editorRights := CheckEditorUser(db, request.EditorId, request.GroupId, w)
	if !editorRights {
		return nil
	}
	_, err = db.Exec("INSERT INTO playergroup (groupid, playerid, groupname, playername, editor) VALUES (?, ?, ?, ?, ?);", strconv.Itoa(request.GroupId), strconv.Itoa(request.PlayerId), request.GroupName, request.PlayerName, 0)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Player entered into group")
	return nil

}

// /groups/changename
func ChangeGroupName(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.ChangeNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	editorRights := CheckEditorUser(db, request.EditorId, request.GroupId, w)
	if !editorRights {
		return err
	}

	_, err = db.Exec("UPDATE groups SET name = ? WHERE groupid = ?;", request.Name, request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	_, err = db.Exec("UPDATE playergroup SET name = ? WHERE groupid = ?;", request.Name, request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Group Name changed to "+request.Name)
	return nil
}

// /groups/updatelastcompleted
func UpdateLastCompleted(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.UpdateLastCompletedGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	_, err = db.Exec("UPDATE groups SET gamelastcompleted = ? WHERE groupid = ?;", request.NewDate, request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Group Last Completed Date changed to "+request.NewDate.String())
	return nil
}

// /groups/giveplayereditor
func GivePlayerEditor(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.GivePlayerEditorRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	editorRights := CheckEditorUser(db, request.CurrentEditorPlayerId, request.GroupId, w)
	if !editorRights {
		return err
	}
	_, err = db.Exec("UPDATE playergroup SET editor = ? WHERE playerid = ? AND groupid = ?;", 1, request.NewEditorPlayerId, request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	return nil
}

// /groups/removeplayereditor
func RemovePlayerEditor(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.GivePlayerEditorRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	editorRights := CheckEditorUser(db, request.CurrentEditorPlayerId, request.GroupId, w)
	if !editorRights {
		return err
	}
	_, err = db.Exec("UPDATE playergroup SET editor = ? WHERE playerid = ? AND groupid = ?;", 0, request.NewEditorPlayerId, request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	return nil
}

// /groups/removeplayer
func RemovePlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.RemovePlayerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	editorRights := CheckEditorUser(db, request.EditorId, request.GroupId, w)
	if !editorRights {
		return err
	}

	_, err = db.Exec("DELETE from playergroup WHERE playerid = ? AND groupid = ?;", request.PlayerToBeRemovedId, request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Player %d is no longer in Group %d by Player %d. ", request.PlayerToBeRemovedId, request.GroupId, request.EditorId)
	return nil
}

// /groups/deletegroup
func DeleteGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.DeleteGroupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	editorRights := CheckEditorUser(db, request.EditorId, request.GroupId, w)
	if !editorRights {
		return err
	}

	_, err = db.Exec("DELETE from groups WHERE groupid = ?;", request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	_, err = db.Exec("DELETE from playergroup WHERE groupid = ?;", request.GroupId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Group #%d is deleted by Player #%d.  ", request.GroupId, request.EditorId)
	return nil
}

func CheckEditorUser(db *sql.DB, playerid int, groupid int, w http.ResponseWriter) (editor bool) {
	fmt.Fprintf(w, "CHECKING PERMISSIONS for %d\n", playerid)
	var editorPermission dto.PlayerGroup
	row := db.QueryRow("SELECT * from playergroup where playerid = ? AND groupid = ?;", playerid, groupid)
	err := row.Scan(&editorPermission.PlayerId, &editorPermission.PlayerName, &editorPermission.GroupId, &editorPermission.GroupName, &editorPermission.Editor)
	editorBool := int(editorPermission.Editor[0])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, err.Error(), 500)
		return
	}
	if editorBool == 1 {
		fmt.Fprintf(w, "Rights approved.")
		return true
	} else {
		fmt.Fprintf(w, "User does not have editor rights.")
		return false
	}
}
