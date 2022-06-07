package games

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

// /games/get
// GET gets a specific game.
func GetGame(db *sql.DB, w http.ResponseWriter, r *http.Request) (*dto.Game, error) {
	var request dto.GetGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM game where gameid = ?;", strconv.Itoa(request.GameId))
	game := new(dto.Game)
	err = row.Scan(&game.GameId, &game.GroupId, &game.Name, &game.Date, &game.TeamOneName, &game.TeamTwoName, &game.TeamOneScore, &game.TeamTwoScore, &game.TeamOneServing, &game.LastUpdated, &game.Completed)
	if utils.Error500Check(err, w) {
		return nil, err
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return game, nil
}

// /games/
// GET gets all games.
func GetAllGames(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.Game, error) {
	rows, err := db.Query("SELECT * FROM game;")
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()

	games := make([]*dto.Game, 0)
	for rows.Next() {
		game := new(dto.Game)
		err := rows.Scan(&game.GameId, &game.GroupId, &game.Name, &game.Date, &game.TeamOneName, &game.TeamTwoName, &game.TeamOneScore, &game.TeamTwoScore, &game.TeamOneServing, &game.LastUpdated, &game.Completed)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		games = append(games, game)
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return games, nil
}

// /games/players
// GET All players in a game.
func GetPlayersInGame(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.PlayerShort, error) {
	var request dto.GetAllPlayersInGame
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	rows, err := db.Query("SELECT playerid, playername FROM playerstatistics where gameid = ?;", strconv.Itoa(request.GameId))
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
	return players, nil
}

// /games/playergames
// GET Get all games for a player.
func GetGamesForPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.GameShort, error) {
	var request dto.GetAllGamesOfPlayer
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	rows, err := db.Query("SELECT gameid, gamename FROM playerstatistics where playerid = ?;", strconv.Itoa(request.PlayerId))
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()
	games := make([]*dto.GameShort, 0)
	for rows.Next() {
		game := new(dto.GameShort)
		err := rows.Scan(&game.GameId, &game.GameName)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		games = append(games, game)
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return games, nil
}

// /games/create
// POST Creates a game.
func CreateGame(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.CreateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	result, err := db.Exec("INSERT INTO game (groupid, name, teamonename, teamtwoname) VALUES (?, ?, ?, ?);", request.GroupId, request.Name, request.TeamOneName, request.TeamTwoName)
	if utils.Error500Check(err, w) {
		return err
	}
	id, err := result.LastInsertId()
	if utils.Error500Check(err, w) {
		return err
	}
	if request.EditorTeam != dto.EditorNotPlaying {
		_, err = db.Exec("INSERT INTO playerstatistics (gameid, playerid, onteamone, playername, gamename) VALUES (?, ?, ?, ?, ?)", id, request.EditorId, request.EditorTeam, request.EditorName, request.Name)
		if utils.Error500Check(err, w) {
			return err
		}
		fmt.Fprintf(w, "Editor added as player. ")
	}
	fmt.Fprintf(w, "Game Created")
	return nil
}

// /games/addplayer
// POST Add player to a game.
func AddPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.AddPlayerToGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return err
	}
	_, err = db.Exec("INSERT INTO playerstatistics (gameid, playerid, onteamone, playername, gamename) VALUES (?, ?, ?, ?, ?)", request.GameId, request.PlayerId, request.OnTeamOne, request.PlayerName, request.GameName)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Player Added \n")
	return nil
}

// /games/changename
// POST Change a game's name.
func ChangeName(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.GameChangeNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}

	_, err = db.Exec("UPDATE game SET name = ? where gameid = ?", request.NewName, request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	_, err = db.Exec("UPDATE playerstatistics SET gamename = ? where gameid = ?", request.NewName, request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Name Changed to %s \n", request.NewName)
	return nil
}

// /games/teamonescore
//	POST Update team one's score.
func UpdateTeamOneScore(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.TeamOneScoreRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return err
	}
	row := db.QueryRow("SELECT teamonescore FROM game where gameid = ?;", request.GameId)
	var teamonescore int
	err = row.Scan(&teamonescore)
	fmt.Fprintf(w, "%d", teamonescore)
	if utils.Error500Check(err, w) {
		return err
	}
	teamonescore = teamonescore + 1
	_, err = db.Exec("UPDATE game SET teamonescore = ? where gameid = ?", teamonescore, request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Team 1 Score Changed To: %d", teamonescore)
	return nil
}

// /games/teamtwoscore
//	POST Update team two's score.
func UpdateTeamTwoScore(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.TeamTwoScoreRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return err
	}
	row := db.QueryRow("SELECT teamtwoscore FROM game where gameid = ?;", request.GameId)
	var teamtwoscore int
	err = row.Scan(&teamtwoscore)
	fmt.Fprintf(w, "%d", teamtwoscore)
	if utils.Error500Check(err, w) {
		return err
	}
	teamtwoscore = teamtwoscore + 1
	_, err = db.Exec("UPDATE game SET teamtwoscore = ? where gameid = ?", teamtwoscore, request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Team 2 Score Changed To: %d", teamtwoscore)
	return nil
}

// /games/completegame
// POST Marks game as complete. No longer can be edited.
func CompleteGame(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.GameCompletedRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	_, err = db.Exec("UPDATE game SET completed = 1 where gameid = ?", request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Game %d has been completed", request.GameId)
	return nil
}

// /games/uncompletegame
// POST Marks game as incomplete. Can now be edited.
func UncompleteGame(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.GameCompletedRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	_, err = db.Exec("UPDATE game SET completed = 0 where gameid = ?", request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Game %d has been set to incomplete.", request.GameId)
	return nil
}

// /games/switchserver
// POST Updates who is serving.
func SwitchServer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.GameSwitchServer
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return err
	}
	if !request.TeamOneServing {
		_, err = db.Exec("UPDATE game SET teamoneserving = 0 where gameid = ?", request.GameId)
		fmt.Fprintf(w, "Team 1 is now serving.")
	} else {
		_, err = db.Exec("UPDATE game SET teamoneserving = 1 where gameid = ?", request.GameId)
		fmt.Fprintf(w, "Team 2 is now serving.")
	}
	if utils.Error500Check(err, w) {
		return err
	}
	return nil
}

// /games/update
// POST Updates the entire game row.
func UpdateGame(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.UpdateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return err
	}
	_, err = db.Exec("UPDATE game SET name = ?, teamonename = ?, teamtwoname = ?, teamonescore = ?, teamtwoscore = ?, teamoneserving = ?, groupid = ?  where gameid = ?", request.Name, request.TeamOneName, request.TeamTwoName, request.TeamOneScore, request.TeamTwoScore, request.TeamOneServing, request.GroupId, request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Game Updated")
	return nil
}

func CheckIfGameIsOpen(db *sql.DB, w http.ResponseWriter, gameid int) bool {
	fmt.Fprintf(w, "CHECKING GAME IS OPEN %d\n", gameid)
	var completed []uint8
	row := db.QueryRow("SELECT completed from game where gameid = ?;", gameid)
	err := row.Scan(&completed)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, err.Error(), 500)
		return false
	}
	completedBool := int(completed[0])

	if completedBool == 0 {
		fmt.Fprintf(w, "Game open.")
		return true
	} else {
		fmt.Fprintf(w, "Game closed.")
		return false
	}
}
