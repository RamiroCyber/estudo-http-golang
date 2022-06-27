package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

func userId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:123456/httpgo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user User

	db.QueryRow("select id,nome from users where id = ? ", id).Scan(&user.ID, &user.Nome)

	json, _ := json2.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456/httpgo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Nome)
		users = append(users, user)
	}

	json, _ := json2.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userId(w, r, id)
	case r.Method == "GET":
		allUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}
