package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Group struct {
	groupid           int
	name              string
	datecreated       time.Time
	gamelastcompleted time.Time
}

var db *sql.DB

func GetDatabase() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println(os.Getenv(("DSN")))
	dsnString := os.Getenv("DSN")
	db, err = sql.Open("mysql", dsnString)
	return db, err
}

func main() {
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
	http.HandleFunc("/", handler)
	http.HandleFunc("/groups", groupsHandler)
	fmt.Println("Listening...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func groupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM groups")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	groups := make([]*Group, 0)
	for rows.Next() {
		group := new(Group)
		err := rows.Scan(&group.groupid, &group.name, &group.datecreated, &group.gamelastcompleted)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, group := range groups {
		fmt.Fprintf(w, "Group ID: %d, Name: %s, Date Created: %s, Game Last Completed: %s\n", group.groupid, group.name, group.datecreated.Format("2006-01-02"), group.gamelastcompleted.Format("2006-01-02"))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

}
