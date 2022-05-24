package groups

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"practice/dto"
	"strconv"
)

func GetAllGroups(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM groups")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, "12", 500)
		return
	}
	defer rows.Close()

	groups := make([]*dto.Group, 0)
	for rows.Next() {
		group := new(dto.Group)
		err := rows.Scan(&group.GroupId, &group.Name, &group.DateCreated, &group.GameLastCompleted)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			http.Error(w, err.Error(), 500)
			return
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, "36", 500)
		return
	}
	for _, group := range groups {
		fmt.Fprintf(w, "Group ID: %d, Name: %s, Date Created: %s, Game Last Completed: %s\n", group.GroupId, group.Name, group.DateCreated.Format("2006-01-02"), group.GameLastCompleted.Format("2006-01-02"))
	}
}

func GetGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var err error
	id := r.URL.Query().Get("groupid")
	var group dto.Group
	query := "SELECT * FROM groups WHERE groupid = " + id
	row := db.QueryRow(query)
	err = row.Scan(&group.GroupId, &group.Name, &group.DateCreated, &group.GameLastCompleted)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "\nGroup ID: %d, Name: %s, Date Created: %s, Game Last Completed: %s\n", group.GroupId, group.Name, group.DateCreated.Format("2006-01-02"), group.GameLastCompleted.Format("2006-01-02"))
}

func GetAllGroupsOfPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GetAllGroups
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "SELECT groupid, groupname FROM playergroup where playerid = " + strconv.Itoa(request.PlayerId)

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		http.Error(w, "12", 500)
		return
	}
	defer rows.Close()

	groups := make([]*dto.GroupShort, 0)
	for rows.Next() {
		group := new(dto.GroupShort)
		err := rows.Scan(&group.GroupId, &group.Name)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			http.Error(w, err.Error(), 500)
			return
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, "36", 500)
		return
	}
	for _, group := range groups {
		fmt.Fprintf(w, "Group ID: %d, Name: %s\n", group.GroupId, group.Name)
	}
}

func CreateGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}

func AddPlayerToGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.AddPlayerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Request: %v\n", request)
	editorRights := checkEditorUser(db, request.EditorId, w)
	if !editorRights {
		return
	}
	// NEED TO CHECK IF PLAYERGROUP EDITORUSER is true
	query := "INSERT INTO playergroup (groupid, playerid, groupname, playername, editor) VALUES (" + strconv.Itoa(request.GroupId) + ", " + strconv.Itoa(request.PlayerId) + ", \"" + request.GroupName + "\", \"" + request.PlayerName + "\", 0);"

	_, err = db.Exec(query)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprintf(w, "Player entered into group")
}

func ChangeGroupName(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
func UpdateLastCompleted(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
func GivePlayerEditor(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
func RemovePlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
func DeleteGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
func checkEditorUser(db *sql.DB, playerid int, w http.ResponseWriter) (editor bool) {
	fmt.Fprintf(w, "CHECKING PERMISSIONS\n")
	query := "SELECT * from playergroup where playerid = " + strconv.Itoa(playerid)
	var editorPermission dto.PlayerGroup
	row := db.QueryRow(query)
	err := row.Scan(&editorPermission.PlayerId, &editorPermission.PlayerName, &editorPermission.GroupId, &editorPermission.Editor, &editorPermission.GroupName)
	fmt.Fprintln(w, editorPermission)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, err.Error(), 500)
		return
	}
	if editorPermission.Editor == 1 {
		return true
	} else {
		fmt.Fprintf(w, "User does not have editor rights.")
		return false
	}
}
