package groups

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Group struct {
	groupid           int
	name              string
	datecreated       time.Time
	gamelastcompleted time.Time
}

func GetAllGroups(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM groups")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, "12", 500)
		return
	}
	defer rows.Close()

	groups := make([]*Group, 0)
	for rows.Next() {
		group := new(Group)
		err := rows.Scan(&group.groupid, &group.name, &group.datecreated, &group.gamelastcompleted)
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
		fmt.Fprintf(w, "Group ID: %d, Name: %s, Date Created: %s, Game Last Completed: %s\n", group.groupid, group.name, group.datecreated.Format("2006-01-02"), group.gamelastcompleted.Format("2006-01-02"))
	}
}

func GetAllGroupsOfPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
}

func AddPlayerToGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	groupid := r.URL.Query().Get("groupid")
	playerid := r.URL.Query().Get("playerid")

	_, err := db.Exec("insert into playergroup (groupid, playerid) VALUES (%d, %d);", groupid, playerid)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprintf(w, "Player entered into group")
}
