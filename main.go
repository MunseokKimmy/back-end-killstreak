package main

import (
	"database/sql"
	"encoding/json"
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
	fmt.Println("Listening...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func groupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/groups/player") {
		//GET QUERY PARAM PLAYERID
		//CALL GROUPS package function
		fmt.Fprintln(w, "Playerrrrr")
		groups.GetAllGroupsOfPlayer(db, w, r)
	} else {
		groups.GetAllGroups(db, w, r)
	}

}

type PersonIDRequest struct {
	id int
}

func handler(w http.ResponseWriter, r *http.Request) {
	var request PersonIDRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Person ID: %d", request.id)

}
