package main

import (
	"crypto/rand"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CheckUser(user User) (bool, string) {

	database, _ := sql.Open("sqlite3", "./test.db")
	//statement, _ := database.Prepare("INSERT INTO Users VALUES (?,?,?,?)")
	//statement.Exec("3", "deneme", "deneme", "deneme")
	query := " SELECT * FROM Users WHERE username='" + user.Username + "' AND password='" + user.Password + "'"
	rows, _ := database.Query(query)
	defer rows.Close()
	var id int
	var compareUser User
	/*var username string
	var password string
	var role string*/
	for rows.Next() {
		rows.Scan(&id, &compareUser.Username, &compareUser.Password, &compareUser.Role)
		user.Role = compareUser.Role
		if user == compareUser {
			return true, user.Role
		}
	}
	return false, ""

}

func AddUser(user User) {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("INSERT INTO Users VALUES (?,?,?,?)")
	RandomCrypto, _ := rand.Prime(rand.Reader, 16)
	statement.Exec(RandomCrypto.String(), user.Username, user.Password, user.Role)

}
