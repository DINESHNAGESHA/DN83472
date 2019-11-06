package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ReadAPI struct {
	Dbsrv   string `json:"DBServer"`
	Dbname  string `json:"DataBase"`
	Dbtbl   string `json:"DatabaseTable"`
	Reffreq string `json:"RefreshFrequency"`
	Reftime string `json:"RefreshTimeEST"`
	Refdur  string `json:"RefreshDurationMin"`
}

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<dbname>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/dbread", getPosts).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var readapi []ReadAPI
	result, err := db.Query("SELECT g.server AS 'DB Server', d.DataBase, d.DatabaseTable, d.RefreshFrequency, d.RefreshTimeEST, d.RefreshDurationMin FROM monkeywrench.DataRefreshSchedule d LEFT JOIN vp_prod.global_settings_db_connections g USING(global_settings_db_connections_idx")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&ReadAPI.Dbsrv, &ReadAPI.Dbname, &ReadAPI.Dbtbl, &ReadAPI.Reffreq, &ReadAPI.Reftime, &ReadAPI.Refdur)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
