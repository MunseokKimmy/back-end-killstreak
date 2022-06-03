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
func GetGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GetGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	row := db.QueryRow("SELECT * FROM game where gameid = ?;", strconv.Itoa(request.GameId))
	game := new(dto.Game)
	err = row.Scan(&game.GameId, &game.GroupId, &game.Name, &game.Date, &game.TeamOneName, &game.TeamTwoName, &game.TeamOneScore, &game.TeamTwoScore, &game.TeamOneServing, &game.LastUpdated, &game.Completed)
	if utils.Error500Check(err, w) {
		return
	}
	if utils.Error500Check(err, w) {
		return
	}
	fmt.Fprintf(w, "Game ID: %d, Name: %s Team 1: %s %d Team 2: %s %d\n ", game.GameId, game.Name, game.TeamOneName, game.TeamOneScore, game.TeamTwoName, game.TeamTwoScore)
}

// /games/
// GET gets all games.
func GetAllGames(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM game;")
	if utils.Error500Check(err, w) {
		return
	}
	defer rows.Close()

	games := make([]*dto.Game, 0)
	for rows.Next() {
		game := new(dto.Game)
		err := rows.Scan(&game.GameId, &game.GroupId, &game.Name, &game.Date, &game.TeamOneName, &game.TeamTwoName, &game.TeamOneScore, &game.TeamTwoScore, &game.TeamOneServing, &game.LastUpdated, &game.Completed)
		if utils.Error500Check(err, w) {
			return
		}
		games = append(games, game)
	}
	if utils.Error500Check(err, w) {
		return
	}
	for _, game := range games {
		fmt.Fprintf(w, "Game ID: %d, Name: %s Team 1: %s %d Team 2: %s %d\n ", game.GameId, game.Name, game.TeamOneName, game.TeamOneScore, game.TeamTwoName, game.TeamTwoScore)
	}
}

// /games/players
// GET All players in a game.
func GetPlayersInGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GetAllPlayersInGame
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	rows, err := db.Query("SELECT playerid, playername FROM playerstatistics where gameid = ?;", strconv.Itoa(request.GameId))
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

// /games/playergames
// GET Get all games for a player.
func GetGamesForPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GetAllGamesOfPlayer
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	rows, err := db.Query("SELECT gameid, gamename FROM playerstatistics where playerid = ?;", strconv.Itoa(request.PlayerId))
	if utils.Error500Check(err, w) {
		return
	}
	defer rows.Close()
	games := make([]*dto.GameShort, 0)
	for rows.Next() {
		game := new(dto.GameShort)
		err := rows.Scan(&game.GameId, &game.GameName)
		if utils.Error500Check(err, w) {
			return
		}
		games = append(games, game)
	}
	if utils.Error500Check(err, w) {
		return
	}
	for _, game := range games {
		fmt.Fprintf(w, "Game ID: %d, Game Name: %s\n", game.GameId, game.GameName)
	}
}

// /games/create
// POST Creates a game.
func CreateGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.CreateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
	result, err := db.Exec("INSERT INTO game (groupid, name, teamonename, teamtwoname) VALUES (?, ?, ?, ?);", request.GroupId, request.Name, request.TeamOneName, request.TeamTwoName)
	if utils.Error500Check(err, w) {
		return
	}
	id, err := result.LastInsertId()
	if utils.Error500Check(err, w) {
		return
	}
	if request.EditorTeam != dto.EditorNotPlaying {
		_, err = db.Exec("INSERT INTO playerstatistics (gameid, playerid, onteamone, playername, gamename) VALUES (?, ?, ?, ?, ?)", id, request.EditorId, request.EditorTeam, request.EditorName, request.Name)
		if utils.Error500Check(err, w) {
			return
		}
		fmt.Fprintf(w, "Editor added as player. ")
	}
	fmt.Fprintf(w, "Game Created")
}

// /games/addplayer
// POST Add player to a game.
func AddPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.AddPlayerToGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return
	}
	_, err = db.Exec("INSERT INTO playerstatistics (gameid, playerid, onteamone, playername, gamename) VALUES (?, ?, ?, ?, ?)", request.GameId, request.PlayerId, request.OnTeamOne, request.PlayerName, request.GameName)
	if utils.Error500Check(err, w) {
		return
	}
	fmt.Fprintf(w, "Player Added \n")

}

// /games/changename
// POST Change a game's name.
func ChangeName(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GameChangeNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}

	_, err = db.Exec("UPDATE game SET name = ? where gameid = ?", request.NewName, request.GameId)
	if utils.Error500Check(err, w) {
		return
	}
	fmt.Fprintf(w, "Name Changed to %s \n", request.NewName)
}

// /games/teamonescore
//	POST Update team one's score.
func UpdateTeamOneScore(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.TeamOneScoreRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return
	}
	row := db.QueryRow("SELECT teamonescore FROM game where gameid = ?;", request.GameId)
	var teamonescore int
	err = row.Scan(&teamonescore)
	fmt.Fprintf(w, "%d", teamonescore)
	if utils.Error500Check(err, w) {
		return
	}
	teamonescore = teamonescore + 1
	_, err = db.Exec("UPDATE game SET teamonescore = ? where gameid = ?", teamonescore, request.GameId)
	if utils.Error500Check(err, w) {
		return
	}
	fmt.Fprintf(w, "Team 1 Score Changed To: %d", teamonescore)
}

// /games/teamtwoscore
//	POST Update team two's score.
func UpdateTeamTwoScore(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.TeamTwoScoreRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
	if !CheckIfGameIsOpen(db, w, request.GameId) {
		return
	}
	row := db.QueryRow("SELECT teamtwoscore FROM game where gameid = ?;", request.GameId)
	var teamtwoscore int
	err = row.Scan(&teamtwoscore)
	fmt.Fprintf(w, "%d", teamtwoscore)
	if utils.Error500Check(err, w) {
		return
	}
	teamtwoscore = teamtwoscore + 1
	_, err = db.Exec("UPDATE game SET teamtwoscore = ? where gameid = ?", teamtwoscore, request.GameId)
	if utils.Error500Check(err, w) {
		return
	}
	fmt.Fprintf(w, "Team 2 Score Changed To: %d", teamtwoscore)
}

// /games/completegame
// POST Marks game as complete. No longer can be edited.
func CompleteGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GameCompletedRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
	_, err = db.Exec("UPDATE game SET completed = 1 where gameid = ?", request.GameId)
	if utils.Error500Check(err, w) {
		return
	}
	fmt.Fprintf(w, "Game %d has been completed", request.GameId)
}

// /games/uncompletegame
// POST Marks game as incomplete. Can now be edited.
func UncompleteGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GameCompletedRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
}

// /games/switchserver
// POST Updates who is serving.
func SwitchServer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GameSwitchServer
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
}

// /games/update
// POST Updates the entire game row.
func UpdateGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return
	}
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
