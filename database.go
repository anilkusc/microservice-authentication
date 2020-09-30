package main

import (
	"crypto/rand"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func CheckUser(user User) (bool, string) {

	database, _ := sql.Open("sqlite3", "./test.db")

	//query := " SELECT * FROM Users WHERE username='" + user.Username + "' AND password='" + user.Password + "'"
	query := " SELECT * FROM Users WHERE username='" + user.Username + "'"
	rows, _ := database.Query(query)
	defer rows.Close()
	var id int
	var compareUser User
	for rows.Next() {
		rows.Scan(&id, &compareUser.Username, &compareUser.Password, &compareUser.Role)
		user.Role = compareUser.Role
		if user.Username == compareUser.Username {
			//compare 2 encrypted passwords
			if err := bcrypt.CompareHashAndPassword([]byte(compareUser.Password), []byte(user.Password)); err != nil {
				return false, ""
			} else {
				return true, user.Role
			}
		}
	}
	return false, ""

}

func AddUser(user User) {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("INSERT INTO Users VALUES (?,?,?,?)")
	RandomCrypto, _ := rand.Prime(rand.Reader, 16)
	//encrypt password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	statement.Exec(RandomCrypto.String(), user.Username, string(hashedPassword), user.Role)

}
