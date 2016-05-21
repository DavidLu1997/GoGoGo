package api

import (
	"html/template"
	"net/http"
	"io"
	"log"
	"encoding/json"
	"database/sql"
	"gogogo/model"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

func ServerStart() {
	http.HandleFunc("/game/", gameHandler)
	http.HandleFunc("/move/", moveHandler)
	http.ListenAndServe(":8070", nil)
}

//Handler to load game
func gameHandler(w http.ResponseWriter, r *http.Request) {
	//Find gameID in database
	id := r.URL.Path[len("/game/"):len("/game/")+idLen]
	b := loadGame(id)
	//Send game via JSON
}

//Handler for moves
func moveHandler(w http.ResponseWriter, r *http.Request) {
	//Handle moves as needed
	idx := len("/move/") + idLen
	id := r.URL.Path[len("/move/"):len("/move/")+idLen]
	player := r.URL.Path[idx:idx+2]
	x := r.URL.Path[idx+2:idx+4]
	y := r.URL.Path[idx+4:idx+6]
	//Process move
	//Write move to DB
	//Rely on client to refresh view
}

//Handler to create new game
func newGameHandler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[len("/game/"):len("/game/")+2]
	size, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	id := randID()

	db, err := sql.Open("sqlite3", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM games WHERE id = ?;", id)
	if err != nil {
		log.Fatal(err)
	}
	while(rows != nil) {
		id = randId()
		ros, err := db.Query("SELECT * FROM games WHERE id = ?;", id)
		if err != nil {
			log.Fatal(err)
		}
	}
	initGame(db, id, size)

	w.Header().Set("Content-Type", "application/javascript")
	idJson := make(map[string]string)
	idJson["id"] = id;
	json.NewEncoder(w).Encode(Payload{idJson})
}

//Initialize game to DB
func initGame(db *sql.DB, id string, size int) {
	board := int[][]
	turn := 0

	//Open DB connection
	db, err := sql.Open("sqlite3", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	//Insert into DB
	res, err := db.Exec("INSERT INTO games(id, board, size, turn) VALUES(?, ?, ?, ?);", id, board, size, turn)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}