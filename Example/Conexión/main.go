package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-oci8"
)

type Tabla struct {
	NombreTabla string `json:"nombreTabla"`
	TipoTabla   string `json:"tipoTabla"`
}

func main() {
	println("start")
	// db, err := sql.Open("oci8", "candido/lord@localhost:1521/xe?PROTOCAL=TCP")
	db, err := sql.Open("oci8", "system/3741@localhost:1521/xe?PROTOCAL=TCP")
	if err != nil {
		log.Fatal(err)
	}
	println("Connection succcess!!")

	// rows, err := db.Query("SELECT sysdate  FROM dual")
	// rows, err := db.Query("SELECT * from cargos")
	rows, err := db.Query("select * from user_catalog")
	// rows, err := db.Query("INSERT INTO SYSTEM.HOLA (IID, NOMBRE) VALUES ('3', 'ae')")

	if err != nil {
		log.Fatal(err)
	}

	var tablas []Tabla
	var str1 string
	var str2 string

	for rows.Next() {
		if err = rows.Scan(&str1, &str2); err != nil {
			log.Fatalln("error fetching", err)
		}

		dato := Tabla{
			NombreTabla: str1,
			TipoTabla:   str2,
		}

		tablas = append(tablas, dato)
	}

	for _, tarea := range tablas {
		fmt.Println("da ", tarea.NombreTabla)
	}

}
