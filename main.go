package main

import (
	"database/sql"
	"fmt"
	"killstreak/groups"
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
	fmt.Println("Listening...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func groupsHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/groups/player") {
		if !utils.Error405CheckGETMethod(w, r) {
			return
		}
		// GET ALL GroupShorts that a player is in. Requires PlayerID.
		groups.GetAllGroupsOfPlayer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/get") {
		if !utils.Error405CheckGETMethod(w, r) {
			return
		}
		// GET Group with GroupID.
		groups.GetGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/addplayer") {
		if !utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// GET Group with GroupID.
		groups.AddPlayerToGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/create") {
		if !utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Group with group name.
		groups.CreateGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/changename") {
		if !utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Group with group name.
		groups.ChangeGroupName(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/updatelastcompleted") {
		if !utils.Error405CheckPOSTMethod(w, r) {
			return
		}
		// POST Group with group name.
		groups.UpdateLastCompleted(db, w, r)
	} else {
		// GET all groups.
		groups.GetAllGroups(db, w, r)
	}

}

func playerHandler(w http.ResponseWriter, r *http.Request) {
}

func handler(w http.ResponseWriter, r *http.Request) {

}
