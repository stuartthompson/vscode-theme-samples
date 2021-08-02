package main

import (
	"log"
)

type User struct {
	UserID       string
}

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * from users")
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	for rows.Next() {
		user := User{}
		scanErr := rows.Scan(&user.UserID, &user.Email, &user.PasswordSalt, &user.PasswordHash, &user.IsActive)
		if scanErr != nil {
			log.Fatal("Error reading row.", scanErr.Error())
		}
		log.Print("Read user: " + user.UserID)
		users = append(users, user)
	}

	log.Print("Success! (I think) :)")
}