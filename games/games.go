package games

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"killstreak/dto"
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
	err = row.Scan(&game.GameId, &game.GroupId, &game.Name, &game.Date, &game.TeamOneName, &game.TeamTwoName, &game.TeamOneScore, &game.TeamTwoScore, &game.TeamOneServing, &game.LastUpdated)
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
		err := rows.Scan(&game.GameId, &game.GroupId, &game.Name, &game.Date, &game.TeamOneName, &game.TeamTwoName, &game.TeamOneScore, &game.TeamTwoScore, &game.TeamOneServing, &game.LastUpdated)
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
	var request dto.GetAllPlayersInGroupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
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

}

// /games/create
// POST Creates a game.
func CreateGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.CreateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
}

// /games/addplayer
// POST Add player to a game.
func AddPlayer(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.AddPlayerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
}

// /games/changename
// POST Change a game's name.
func ChangeName(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GameChangeNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
}

// /games/teamonescore
//	POST Update team one's score.
func UpdateTeamOneScore(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.TeamOneScoreRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
}

// /games/teamtwoscore
//	POST Update team two's score.
func UpdateTeamTwoScore(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.TeamTwoScoreRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
}

// /games/completegame
// POST Marks game as complete. No longer can be edited.
func CompleteGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GameCompletedRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
}

// /games/uncompletegame
// POST Marks game as incomplete. Can now be edited.
func UncompleteGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GameCompletedRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
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
}

// /games/update
// POST Updates the entire game row.
func UpdateGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
}
