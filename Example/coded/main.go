package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde Get")
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde Post")
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde Put")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde Delete")
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/user", GetUsers).Methods("GET")
	r.HandleFunc("/api/user", PostUsers).Methods("POST")
	r.HandleFunc("/api/user", PutUsers).Methods("PUT")
	r.HandleFunc("/api/user", DeleteUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Escuchando...")
	server.ListenAndServe()
}
