package forum

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Create(phrase string, user string, title string, categorie string) string {
	if phrase == "" || user == "" || title == "" || categorie == "" {
		return "Empty fields"
	} else {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		defer db.Close()
		stmt, err := db.Prepare("INSERT INTO Topics (title, message, date, categorie, creatorname,like, dislike) VALUES (?, ?, ?, ?, ?, ?, ?)")
		CheckErr(err)
		defer stmt.Close()
		time := time.Now().Format("2 Jan 2006 15:04")
		_, err = stmt.Exec(title, phrase, time, categorie, user, 0, 0)
		CheckErr(err)
		return "Topic created"
	}
}
