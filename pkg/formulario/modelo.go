package formulario

import (
	"log"

	"github.com/test/pkg/db"
	"github.com/test/pkg/res"
)

func CreateTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS formulario (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT
		);`

	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table formulario created")

	return nil
}

func InsertFormulario(name, descripcion string) error {
	sql := `
		INSERT INTO formulario (name, description)
		VALUES (?, ?);`

	_, err := db.DB.Exec(sql, name, descripcion)
	if err != nil {
		return err
	}

	log.Println("Formulario inserted")

	return nil
}

func SelectAllFormularios() (res.Json, error) {
	sql := `
		SELECT * FROM formulario;`

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var formulario []res.Json
	for rows.Next() {
		var id int
		var name string
		var description string
		

		err = rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}

		formulario = append(formulario, res.Json{
			"id":       id,
			"name":     name,
			"description":    description,
			
		})
	}

	return res.Json{
		"formulario": formulario,
	}, nil
}

func Updateformulario(id int, name, descripcion string) error {
	sql := `
		UPDATE formulario SET
		name = ?,
		description = ?
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, name, descripcion, id)
	if err != nil {
		return err
	}

	log.Println("formulario updated")

	return nil
}

func DeleteFormulario(id int) error {
	sql := `
		DELETE FROM formulario
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, id)
	if err != nil {
		return err
	}

	log.Println("formulario deleted")

	return nil
}
