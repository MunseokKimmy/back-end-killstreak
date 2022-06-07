package players

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"killstreak/dto"
	"killstreak/groups"
	"killstreak/utils"
	"net/http"
	"strconv"
)

// /player/getplayersingroup
func GetAllPlayersInGroup(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.PlayerShort, error) {
	var request dto.GetAllPlayersInGroupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	rows, err := db.Query("SELECT playerid, playername FROM playergroup where groupid = ?;", strconv.Itoa(request.GroupId))
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()

	players := make([]*dto.PlayerShort, 0)
	for rows.Next() {
		player := new(dto.PlayerShort)
		err := rows.Scan(&player.PlayerId, &player.Name)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		players = append(players, player)
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	for _, player := range players {
		fmt.Fprintf(w, "Player ID: %d, Name: %s\n", player.PlayerId, player.Name)
	}
	return players, nil
}

// /player/get?playerid=1
func GetPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) (*dto.Player, error) {
	id := r.URL.Query().Get("playerid")
	player := new(dto.Player)
	row := db.QueryRow("SELECT * FROM player WHERE playerid = ?;", id)
	err := row.Scan(&player.PlayerId, &player.FirstName, &player.LastName, &player.Kills, &player.Assists, &player.Aces,
		&player.Digs, &player.Blocks, &player.AtkErrors, &player.ServiceErrors, &player.AssistErrors, &player.BlockErrors, &player.AccountId)
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return player, nil
}

// /players/
func GetAllPlayers(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.Player, error) {
	rows, err := db.Query("SELECT * from player;")
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()

	players := make([]*dto.Player, 0)
	for rows.Next() {
		player := new(dto.Player)
		err := rows.Scan(&player.PlayerId, &player.FirstName, &player.LastName, &player.Kills, &player.Assists, &player.Aces,
			&player.Digs, &player.Blocks, &player.AtkErrors, &player.ServiceErrors, &player.AssistErrors, &player.BlockErrors, &player.AccountId)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		players = append(players, player)
	}
	err = rows.Err()
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return players, nil
}

// /player/create
func CreatePlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.CreatePlayerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	result, err := db.Exec("INSERT INTO player (firstname, lastname) values (?, ?);", request.FirstName, request.LastName)
	if utils.Error500Check(err, w) {
		return err
	}
	id, err := result.LastInsertId()
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "New Player Created: %s %s - %d", request.FirstName, request.LastName, id)
	return nil
}

// /player/changename
func ChangePlayerName(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.ChangePlayerNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if request.AccountId == 0 {
		if groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
			fullName := request.FirstName + " " + request.LastName
			_, err := db.Exec("UPDATE player SET firstname = ?, lastname = ? WHERE playerid = ?", request.FirstName, request.LastName, request.PlayerId)
			if err != nil {
				fmt.Fprint(w, err.Error())
				return err
			}
			_, err = db.Exec("UPDATE playerstatistics SET playername = ? WHERE playerid = ?", fullName, request.PlayerId)
			if err != nil {
				fmt.Fprint(w, err.Error())
				return err
			}
			_, err = db.Exec("UPDATE playergroup SET playername = ? WHERE playerid = ?", fullName, request.PlayerId)
			if err != nil {
				fmt.Fprint(w, err.Error())
				return err
			}
			fmt.Fprintf(w, "Player %d name changed to %s %s. ", request.PlayerId, request.FirstName, request.LastName)
		} else {
			return err
		} //TODO: Finish when Account methods are set up.
	} else {

	}
	return nil
}

// /player/updatelifetimetotals
// This will take in the NEW LifetimeTotals, it will not add them together.
// Potentially in the future, we may need to update them in a way that's not adding, such as removing some errors.
func UpdatePlayerLifetimeTotals(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.UpdatePlayerLifetimeTotalsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	_, err = db.Exec("UPDATE player SET kills = ?, assists = ?, aces = ?, digs = ?, blocks = ?, atkerrors = ?, serviceerrors = ?, assisterrors = ?, blockerrors = ? WHERE playerid = ?;", request.Kills, request.Assists, request.Aces, request.Digs, request.Blocks, request.AtkErrors, request.ServiceErrors, request.AssistErrors, request.BlockErrors, request.PlayerId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Player %d stats changed to Kills: %d Assists: %d Aces: %d Digs: %d Blocks: %d. ", request.PlayerId, request.Kills, request.Assists, request.Aces, request.Digs, request.Blocks)
	return nil
}
