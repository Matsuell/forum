package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func EditMail(username string, mail string, mailconf string) {
	if mail == mailconf {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_edit_mail, err := db.Prepare("UPDATE User SET email=? WHERE username=?")
		CheckErr(err)
		request_edit_mail.Exec(mail, username)
		db.Close()
	}
}

func EditPassword(username string, currentpassword string, password string, passwordconf string) {
	if password == passwordconf {
		db, err := sql.Open("sqlite3", filedb)
		CheckErr(err)
		request_current_password := "SELECT password FROM User WHERE username='" + username + "'"
		CheckErr(err)
		var currentpasshash []byte

		rows, err := db.Query(request_current_password)
		CheckErr(err)

		for rows.Next() {
			err = rows.Scan(&currentpasshash)
			CheckErr(err)
		}

		CheckErr(err)
		err = bcrypt.CompareHashAndPassword(currentpasshash, []byte(currentpassword))
		CheckErr(err)
		request_edit_password, err := db.Prepare("UPDATE User SET password=? WHERE username=?")
		CheckErr(err)
		passhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		CheckErr(err)
		request_edit_password.Exec(passhash, username)
		db.Close()
	}
}

func EditUsername(username string, newusername string) {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_edit_username, err := db.Prepare("UPDATE User SET username=? WHERE username=?")
	CheckErr(err)
	request_edit_username.Exec(newusername, username)
	db.Close()
}

func DeleteAccount(username string) {
	db, err := sql.Open("sqlite3", filedb)
	CheckErr(err)
	request_delete_account, err := db.Prepare("DELETE From User WHERE username= ?")
	CheckErr(err)
	request_delete_account.Exec(username)
	db.Close()
}
