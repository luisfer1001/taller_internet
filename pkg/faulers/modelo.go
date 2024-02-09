package faulers

import (
	"log"

	"github.com/test/pkg/db"
	"github.com/test/pkg/res"
)

func CreateTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS faulers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code INTEGER,
			faules TEXT
		);`

	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table faulers created")

	return nil
}

func InsertFauler(code int, faules string) error {
	sql := `
		INSERT INTO faulers (code, faules)
		VALUES (?, ?);`

	_, err := db.DB.Exec(sql, code, faules)
	if err != nil {
		return err
	}

	log.Println("Fauler inserted")

	return nil
}

func SelectAllFaulers() (res.Json, error) {
	sql := `
		SELECT * FROM faulers;`

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faulers []res.Json
	for rows.Next() {
		var id int
		var code int
		var faules string

		err = rows.Scan(&id, &code, &faules)
		if err != nil {
			return nil, err
		}

		faulers = append(faulers, res.Json{
			"id":     id,
			"code":   code,
			"faules": faules,
		})
	}

	return res.Json{
		"faulers": faulers,
	}, nil
}

func UpdateFauler(id, code int, faules string) error {
	sql := `
		UPDATE faulers SET
		code = ?,
		faules = ?,
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, code, faules, id)
	if err != nil {
		return err
	}

	log.Println("Fauler updated")

	return nil
}

func DeleteFauler(id int) error {
	sql := `
		DELETE FROM faulers
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, id)
	if err != nil {
		return err
	}

	log.Println("Fauler deleted")

	return nil
}
