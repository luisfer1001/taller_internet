package imagenes

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
			file TEXT,
			id_formulario INT
		);`

	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Table imagenes created")

	return nil
}

func InsertImagen(name, file string, id_formulario int) error {
	sql := `
		INSERT INTO imagenes (name, file, id_formulario)
		VALUES (?, ?, ?);`

	_, err := db.DB.Exec(sql, name, file, id_formulario)
	if err != nil {
		return err
	}

	log.Println("Imagen inserted")

	return nil
}

func SelectAllImagenes() (res.Json, error) {
	sql := `
		SELECT * FROM imagenes;`

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var imagenes []res.Json
	for rows.Next() {
		var id int
		var name string
		var file string
		var id_formulario int

		err = rows.Scan(&id, &name, &file, &id_formulario)
		if err != nil {
			return nil, err
		}

		imagenes = append(imagenes, res.Json{
			"id":            id,
			"name":          name,
			"file":          file,
			"id_formulario": id_formulario,
		})
	}

	return res.Json{
		"imagenes": imagenes,
	}, nil
}

func UpdateImagen(id int, name, file string, id_formulario int) error {
	sql := `
		UPDATE imagenes SET
		name = ?,
		file = ?,
		id_formulario = ?
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, name, file, id_formulario, id)
	if err != nil {
		return err
	}

	log.Println("Imagen updated")

	return nil
}

func DeleteImagen(id int) error {
	sql := `
		DELETE FROM imagenes
		WHERE id = ?;`

	_, err := db.DB.Exec(sql, id)
	if err != nil {
		return err
	}

	log.Println("Imagen deleted")

	return nil
}
