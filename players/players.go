package players

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"killstreak/dto"
	"killstreak/utils"
	"net/http"
	"strconv"
)

// /groups/getplayersingroup
func GetAllPlayersInGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GetAllPlayersInGroupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	fmt.Fprintf(w, "Request %v", request)
	rows, err := db.Query("SELECT playerid, playername FROM playergroup where groupid = ?;", strconv.Itoa(request.GroupId))
	if utils.Error500Check(err, w) {
		return
	}
	defer rows.Close()

	players := make([]*dto.PlayerShort, 0)
	for rows.Next() {
		player := new(dto.PlayerShort)
		err := rows.Scan(&player.PlayerId, &player.Name)
		if utils.Error500Check(err, w) {
			return
		}
		players = append(players, player)
	}
	if utils.Error500Check(err, w) {
		return
	}
	for _, player := range players {
		fmt.Fprintf(w, "Player ID: %d, Name: %s\n", player.PlayerId, player.Name)
	}
}
