package stats

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"killstreak/dto"
	"killstreak/utils"
	"net/http"
)

// /stats/getgamestats/
func GetGameStats(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var request dto.GetGameStatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if utils.Error400Check(err, w) {
		return
	}
	rows, err := db.Query("SELECT * from playerstatistics WHERE gameid = ?", request.GameId)
	if utils.Error500Check(err, w) {
		return
	}
	defer rows.Close()

	gameStats := make([]*dto.StatBlock, 0)
	for rows.Next() {
		block := new(dto.StatBlock)
		err := rows.Scan(&block.StatId, &block.GameId, &block.PlayerId, &block.Kills, &block.AtkErrors, &block.ServiceAces, &block.ServiceErrors, &block.Assists, &block.AssistErrors, &block.Digs, &block.Blocks, &block.BlockErrors, &block.OnTeamOne, &block.PlayerName, &block.GameName)
		if utils.Error500Check(err, w) {
			return
		}
		gameStats = append(gameStats, block)
	}
	if utils.Error500Check(err, w) {
		return
	}
	for _, block := range gameStats {
		fmt.Fprintf(w, "Stat ID: %d \n Game ID: %d %s \n Player: %d - %s\n Kills: %d\n AtkErrors: %d\n Aces: %d\n ServiceErrors: %d\n Assists: %d\n AssistErrors: %d\n Digs: %d\n, Blocks: %d\n, Block Errors: %d\n On Team One: %d\n",
			block.StatId, block.GameId, block.GameName, block.PlayerId, block.PlayerName, block.Kills, block.AtkErrors, block.ServiceAces, block.ServiceErrors, block.Assists, block.AssistErrors, block.Digs, block.Blocks, block.BlockErrors, int(block.OnTeamOne[0]))
	}
}

/*
	/stats/getplayerstatsingame/
	Gets all stats for a player in a game.
*/
func GetPlayersStatsInGame(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}

/*
	/stats/getplayerstats/
	Gets all stats for a player.
*/
func GetPlayersStats(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}

/*
	/stats/create/
	Creates a new set of stats. Used when a game is created with a player in it.
	Requires editor user.
*/
func CreateStats(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}

/*
	/stats/update/
	Updates block of stats for a player. Used during the game for any stat recorded.
	Requires editor user.
*/
func UpdateStats(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}

/*
	/stats/highlights/
	Gets all group stats and adds them together.
*/
func Highlights(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}

/*
	/stats/allgroupstats/
	Gets Group Stats On a certain day.
*/
func AllGroupStats(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
