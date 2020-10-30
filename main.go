package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-oci8"
)

func main() {
	println("start")
	db, err := sql.Open("oci8", "system/3741@localhost:1521/xe?PROTOCAL=TCP")
	if err != nil {
		log.Fatal(err)
	}
	println("Connection succcess!!")
	// rows, err := db.Query("SELECT sysdate  FROM dual")
	rows, err := db.Query("SELECT CARGO from cargos")
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
