package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"practice/groups"
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
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/groups/player") {
		// GET ALL GroupShorts that a player is in. Requires PlayerID.
		groups.GetAllGroupsOfPlayer(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/get") {
		// GET Group with GroupID.
		groups.GetGroup(db, w, r)
	} else if strings.HasPrefix(r.URL.Path, "/groups/addplayer") {
		// GET Group with GroupID.
		groups.AddPlayerToGroup(db, w, r)
	} else {
		// GET all groups.
		groups.GetAllGroups(db, w, r)
	}

}

func playerHandler(w http.ResponseWriter, r *http.Request) {
}

func handler(w http.ResponseWriter, r *http.Request) {

}
