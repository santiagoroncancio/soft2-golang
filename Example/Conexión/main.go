package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-oci8"
)

// type Tabla struct {
// 	NombreTabla string `json:"nombreTabla"`
// 	TipoTabla   string `json:"tipoTabla"`
// }

func main() {
	println("start")
	// db, err := sql.Open("oci8", "candido/lord@localhost:1521/xe?PROTOCAL=TCP")
	db, err := sql.Open("oci8", "system/3741@localhost:1521/xe?PROTOCAL=TCP")
	if err != nil {
		log.Fatal(err)
	}
	println("Connection succcess!!")

	// rows, err := db.Query("SELECT sysdate  FROM dual")
	rows, err := db.Query("SELECT * from cargos")
	// rows, err := db.Exec("select column_name from all_tab_columns where table_name = 'HOLA';")
	// rows, err := db.Query("INSERT INTO SYSTEM.HOLA (IID, NOMBRE) VALUES ('3', 'ae')")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows)

	// var str1 string

	// for rows.Next() {
	// 	if err = rows.Scan(&str1); err != nil {
	// 		log.Fatalln("error fetching", err)
	// 	}

	// 	fmt.Println(str1)
	// }

}
