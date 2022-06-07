package stats

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"killstreak/dto"
	"killstreak/games"
	"killstreak/groups"
	"killstreak/utils"
	"net/http"
	"time"
)

// /stats/getgamestats/
func GetGameStats(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.StatBlock, error) {
	var request dto.GetGameStatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	rows, err := db.Query("SELECT * from playerstatistics WHERE gameid = ?", request.GameId)
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()

	gameStats := make([]*dto.StatBlock, 0)
	for rows.Next() {
		block := new(dto.StatBlock)
		err := rows.Scan(&block.StatId, &block.GameId, &block.PlayerId, &block.Kills, &block.AtkErrors, &block.ServiceAces, &block.ServiceErrors, &block.Assists, &block.AssistErrors, &block.Digs, &block.Blocks, &block.BlockErrors, &block.PlayerName, &block.GameName, &block.Team)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		gameStats = append(gameStats, block)
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return gameStats, nil
}

/*
	/stats/getplayerstatsingame/
	Gets all stats for a player in a game.
*/
func GetPlayersStatsInGame(db *sql.DB, w http.ResponseWriter, r *http.Request) (*dto.StatBlock, error) {
	var request dto.GetPlayersStatsInGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	row := db.QueryRow("SELECT * from playerstatistics WHERE gameid = ? AND playerid = ?", request.GameId, request.PlayerId)

	block := new(dto.StatBlock)
	err = row.Scan(&block.StatId, &block.GameId, &block.PlayerId, &block.Kills, &block.AtkErrors, &block.ServiceAces, &block.ServiceErrors, &block.Assists, &block.AssistErrors, &block.Digs, &block.Blocks, &block.BlockErrors, &block.PlayerName, &block.GameName, &block.Team)
	if utils.Error500Check(err, w) {
		return nil, err
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return block, nil
}

/*
	/stats/getplayerstats/
	Gets all stats for a player.
*/
func GetPlayersStats(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.StatBlock, error) {
	var request dto.GetPlayersStatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	rows, err := db.Query("SELECT * from playerstatistics WHERE playerid = ?", request.PlayerId)
	if utils.Error500Check(err, w) {
		return nil, err
	}
	defer rows.Close()

	gameStats := make([]*dto.StatBlock, 0)
	for rows.Next() {
		block := new(dto.StatBlock)
		err := rows.Scan(&block.StatId, &block.GameId, &block.PlayerId, &block.Kills, &block.AtkErrors, &block.ServiceAces, &block.ServiceErrors, &block.Assists, &block.AssistErrors, &block.Digs, &block.Blocks, &block.BlockErrors, &block.PlayerName, &block.GameName, &block.Team)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		gameStats = append(gameStats, block)
	}
	if utils.Error500Check(err, w) {
		return nil, err
	}
	return gameStats, nil
}

/*
	/stats/switchplayerteam/
	Sets a player's team. Requires editor user.
*/
func SetPlayerTeam(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.SetPlayerTeamRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	if !games.CheckIfGameIsOpen(db, w, request.GameId) {
		return err
	}
	_, err = db.Exec("UPDATE playerstatistics SET team = ? WHERE playerid = ? AND gameid = ?;", request.Team, request.PlayerId, request.GameId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Team Updated \n")
	return nil
}

/*
	/stats/update/
	Updates block of stats for a player. Used during the game for any stat recorded.
	Requires editor user.
*/
func UpdateStats(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var request dto.UpdateStatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return err
	}
	if !groups.CheckEditorUser(db, request.EditorId, request.GroupId, w) {
		return err
	}
	if !games.CheckIfGameIsOpen(db, w, request.GameId) {
		return err
	}
	_, err = db.Exec("UPDATE playerstatistics SET kills = ?, atkerrors = ?, serviceaces = ?, serviceerrors = ?, assists = ?, assisterrors = ?, digs = ?, blocks = ?, blockerrors = ? WHERE gameid = ? AND playerid = ?;",
		request.Kills, request.AtkErrors, request.ServiceAces, request.ServiceErrors, request.Assists, request.AssistErrors, request.Digs, request.Blocks, request.BlockErrors, request.GameId, request.PlayerId)
	if utils.Error500Check(err, w) {
		return err
	}
	fmt.Fprintf(w, "Stats updated \n")
	return nil
}

/*
	/stats/highlights/
	Gets all group stats and adds them together.
*/
func Highlights(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.HighlightsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	date, err := time.Parse("2006-01-02", request.Date)
	if utils.Error500Check(err, w) {
		return
	}
	rows, err := db.Query("SELECT gameid, name from game WHERE date between ? and ?", date.Format("2006-01-02"), date.Format("2006-01-02"))
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
	gameStats := make([]*dto.StatBlock, 0)
	for _, game := range games {
		fmt.Fprintf(w, "Game ID: %d, Game Name: %s\n", game.GameId, game.GameName)
		rows, err := db.Query("SELECT * from playerstatistics WHERE gameid = ?", game.GameId)
		if utils.Error500Check(err, w) {
			return
		}
		defer rows.Close()
		for rows.Next() {
			block := new(dto.StatBlock)
			err := rows.Scan(&block.StatId, &block.GameId, &block.PlayerId, &block.Kills, &block.AtkErrors, &block.ServiceAces, &block.ServiceErrors, &block.Assists, &block.AssistErrors, &block.Digs, &block.Blocks, &block.BlockErrors, &block.Team, &block.PlayerName, &block.GameName)
			if utils.Error500Check(err, w) {
				return
			}
			gameStats = append(gameStats, block)
		}
		if utils.Error500Check(err, w) {
			return
		}
	}
	for _, block := range gameStats {
		fmt.Fprintf(w, "Stat ID: %d \n Game ID: %d %s \n Player: %d - %s\n Kills: %d\n AtkErrors: %d\n Aces: %d\n ServiceErrors: %d\n Assists: %d\n AssistErrors: %d\n Digs: %d\n, Blocks: %d\n, Block Errors: %d\n On Team One: %d\n",
			block.StatId, block.GameId, block.GameName, block.PlayerId, block.PlayerName, block.Kills, block.AtkErrors, block.ServiceAces, block.ServiceErrors, block.Assists, block.AssistErrors, block.Digs, block.Blocks, block.BlockErrors, block.Team)
	}
}

/*
	/stats/allgroupstats/
	Gets Group Stats On a certain day.
*/
func AllGroupStats(db *sql.DB, w http.ResponseWriter, r *http.Request) ([]*dto.StatBlock, error) {
	var request dto.AllGroupStatsOnDayRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return nil, err
	}
	date, err := time.Parse("2006-01-02", request.Date)
	if utils.Error500Check(err, w) {
		return nil, err
	}
	rows, err := db.Query("SELECT gameid, name from game WHERE date between ? and ?", date.Format("2006-01-02"), date.Format("2006-01-02"))
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
	gameStats := make([]*dto.StatBlock, 0)
	for _, game := range games {
		rows, err := db.Query("SELECT * from playerstatistics WHERE gameid = ?", game.GameId)
		if utils.Error500Check(err, w) {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			block := new(dto.StatBlock)
			err := rows.Scan(&block.StatId, &block.GameId, &block.PlayerId, &block.Kills, &block.AtkErrors, &block.ServiceAces, &block.ServiceErrors, &block.Assists, &block.AssistErrors, &block.Digs, &block.Blocks, &block.BlockErrors, &block.PlayerName, &block.GameName, &block.Team)
			if utils.Error500Check(err, w) {
				return nil, err
			}
			gameStats = append(gameStats, block)
		}
		if utils.Error500Check(err, w) {
			return nil, err
		}
	}
	return gameStats, nil
}
