package fauresAreas

import (
	"log"

	"github.com/test/pkg/db"
	"github.com/test/pkg/res"
)

func CreateTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS faures_areas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			id_fauler INTEGER,
			name_area TEXT,
		);`

	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table faures_areas created")

	return nil
}

func InsertFaures(idFauler, nameArea string) error {
	sql := `
		INSERT INTO faures_areas (id_fauler, name_area)
		VALUES (?, ?);`

	_, err := db.DB.Exec(sql, idFauler, nameArea)
	if err != nil {
		return err
	}

	log.Println("")

	return nil
}

func selectFauresAreas() (res.Json, error) {
	sql := `
		SELECT * FROM faures_areas;`

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faures []res.Json
	for rows.Next() {
		var id int
		var idFauler string
		var nameArea string

		err = rows.Scan(&id, &idFauler, &nameArea)
		if err != nil {
			return nil, err
		}

		faures = append(faures, res.Json{
			"id":        id,
			"id_fauler": idFauler,
			"name_area": nameArea,
		})
	}

	return res.Json{
		"faures": faures,
	}, nil
}

func UpdateFaures(id int, idFauler, nameArea string) error {
	sql := `
		UPDATE faures_areas SET
		id_fauler = ?,
		name_area = ?,
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, idFauler, nameArea, id)
	if err != nil {
		return err
	}

	log.Println("Faures updated")

	return nil
}

func DeleteFaures(id int) error {
	sql := `
		DELETE FROM faures_areas
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, id)
	if err != nil {
		return err
	}

	log.Println("Faures deleted")

	return nil
}
