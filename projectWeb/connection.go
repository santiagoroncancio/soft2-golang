package main

import "database/sql"

var db *sql.DB

// GetConnection - Funcion de la conexion
func GetConnection(user string, pass string) *sql.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("oci8", user+"/"+pass+"@localhost:1521/xe?PROTOCAL=TCP")
	if err != nil {
		panic(err)
	}
	return db
}

func get2Con() *sql.DB {
	if db != nil {
		return db
	}
	return db
}
