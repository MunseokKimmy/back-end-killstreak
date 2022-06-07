package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"killstreak/games"
	"killstreak/groups"
	"killstreak/players"
	"killstreak/stats"
	"killstreak/utils"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func GetDatabase() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dsnString := os.Getenv("DSN")
	db, err := sql.Open("mysql", dsnString)
	return db, err
}

func main() {
	var err error
	db, err = GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
	http.HandleFunc("/practice/", handler)
	http.HandleFunc("/groups/", groupsHandler)
	http.HandleFunc("/player/", playerHandler)
	http.HandleFunc("/game/", gameHandler)
	http.HandleFunc("/stats/", statsHandler)
	fmt.Println("Listening...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func groupsHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/groups/player") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		// GET ALL GroupShorts that a player is in. Requires PlayerID.
		groups.GetAllGroupsOfPlayer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/get") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		// GET Group with GroupID.
		groups.GetGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/addplayer") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// GET Group with GroupID.
		groups.AddPlayerToGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/create") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Create Group
		groups.CreateGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/changename") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Change group name.
		groups.ChangeGroupName(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/updatelastcompleted") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Update last completed game date.
		groups.UpdateLastCompleted(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/giveplayereditor") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Give Player Editor
		groups.GivePlayerEditor(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/removeplayereditor") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Remove player editor.
		groups.RemovePlayerEditor(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/removeplayer") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Remove player from group.
		groups.RemovePlayer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/deletegroup") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Delete Group.
		groups.DeleteGroup(db, w, r)
	} else {
		// GET all groups.
		groups.GetAllGroups(db, w, r)
	}
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/player/create") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Create Player.
		players.CreatePlayer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/player/changename") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Change player name.
		players.ChangePlayerName(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/player/updatelifetimetotals") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		//POST Update Player's Lifetime totals.
		players.UpdatePlayerLifetimeTotals(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/player/getplayersingroup") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		// GET Get all players in a group.
		players.GetAllPlayersInGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/player/get") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		// GET Get one player.
		players.GetPlayer(db, w, r)
	} else {
		//GET Get all players.
		players.GetAllPlayers(db, w, r)
	}
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/game/get") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		games.GetGame(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/players") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		games.GetPlayersInGame(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/playergames") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		games.GetGamesForPlayer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/create") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.CreateGame(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/addplayer") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.AddPlayer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/changename") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.ChangeName(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/teamonescore") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.UpdateTeamOneScore(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/teamtwoscore") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.UpdateTeamTwoScore(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/completegame") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.CompleteGame(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/uncompletegame") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.UncompleteGame(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/switchserver") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.SwitchServer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/game/update") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		games.UpdateGame(db, w, r)
	} else {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		games.GetAllGames(db, w, r)
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if strings.HasPrefix(r.URL.Path, "/stats/getgamestats") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		stats, err := stats.GetGameStats(db, w, r)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(stats)
		}
	} else if strings.HasPrefix(r.URL.Path, "/stats/getplayerstatsingame") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		stats, err := stats.GetPlayersStatsInGame(db, w, r)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(stats)
		}
	} else if strings.HasPrefix(r.URL.Path, "/stats/getplayerstats") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		stats, err := stats.GetPlayersStats(db, w, r)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(stats)
		}
	} else if strings.HasPrefix(r.URL.Path, "/stats/setteam") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		err := stats.SetPlayerTeam(db, w, r)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
	} else if strings.HasPrefix(r.URL.Path, "/stats/update") {
		if utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		err := stats.UpdateStats(db, w, r)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
	} else if strings.HasPrefix(r.URL.Path, "/stats/highlights") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		stats.Highlights(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/stats/allgroupstats") {
		if utils.Error405CheckGETMethod(w, r) {
			return
		}
		stats, err := stats.AllGroupStats(db, w, r)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(stats)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

}
