package main

import "database/sql"

var db *sql.DB

// GetConnection - Funcion de la conexion
func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("oci8", "system/3741@localhost:1521/xe?PROTOCAL=TCP")
	if err != nil {
		panic(err)
	}
	return db
}
