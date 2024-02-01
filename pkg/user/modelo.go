package user

import (
	"log"

	"github.com/test/pkg/db"
	"github.com/test/pkg/res"
)

func CreateTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			password TXT
		);`

	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table users created")

	return nil
}

func InsertUser(name, email, password string) error {
	sql := `
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?);`

	_, err := db.DB.Exec(sql, name, email, password)
	if err != nil {
		return err
	}

	log.Println("User inserted")

	return nil
}

func SelectAllUsers() (res.Json, error) {
	sql := `
		SELECT * FROM users;`

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []res.Json
	for rows.Next() {
		var id int
		var name string
		var email string
		var password string

		err = rows.Scan(&id, &name, &email, &password)
		if err != nil {
			return nil, err
		}

		users = append(users, res.Json{
			"id":       id,
			"name":     name,
			"email":    email,
			"password": password,
		})
	}

	return res.Json{
		"users": users,
	}, nil
}

func UpdateUser(id int, name, email, password string) error {
	sql := `
		UPDATE users SET
		name = ?,
		email = ?,
		password = ?
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, name, email, password, id)
	if err != nil {
		return err
	}

	log.Println("User updated")

	return nil
}

func DeleteUser(id int) error {
	sql := `
		DELETE FROM users
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, id)
	if err != nil {
		return err
	}

	log.Println("User deleted")

	return nil
}
