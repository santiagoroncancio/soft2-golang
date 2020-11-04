package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-oci8"
)

func main() {
	println("start")
	// db, err := sql.Open("oci8", "candido/lord@localhost:1521/xe?PROTOCAL=TCP")
	db, err := sql.Open("oci8", "system/3741@localhost:1521/xe?PROTOCAL=TCP")
	if err != nil {
		log.Fatal(err)
	}
	println("Connection succcess!!")

	// rows, err := db.Query("SELECT sysdate  FROM dual")
	// rows, err := db.Query("SELECT cargo from cargos")
	rows, err := db.Query("INSERT INTO SYSTEM.HOLA (IID, NOMBRE) VALUES ('3', 'ae')")

	if err != nil {
		log.Fatal(err)
	}
	var (
		sysdate string
	)
	for rows.Next() {
		if err = rows.Scan(&sysdate); err != nil {
			log.Fatalln("error fetching", err)
		}
		log.Println(sysdate)
	}
}
