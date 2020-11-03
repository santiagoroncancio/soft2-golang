package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-oci8"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Note - struct
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}

// User - struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var noteStore = make(map[string]Note)

var id int

// GetNoteHandler - GET - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, value := range noteStore {
		notes = append(notes, value)
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// PostNoteHandler - POST - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedAt = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// PutNoteHandler - PUT - /api/notes
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteUpdate.CreatedAt = note.CreatedAt
		delete(noteStore, k)
		noteStore[k] = noteUpdate
	} else {
		log.Printf("No se encontro el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteNoteHandler - DELETE - /api/notes
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("No se encontro el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

// LoginUser - Inicio de sesion y conexion con oracle
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	println("start")
	db, err := sql.Open("oci8", user.Username+"/"+user.Password+"@localhost:1521/xe?PROTOCAL=TCP")
	if err != nil {
		log.Fatal(err)
	}
	println("Connection succcess!!")
	rows, err := db.Query("SELECT sysdate  FROM dual")
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

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	// Cors
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/api/login", LoginUser).Methods("POST")

	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	// server := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// log.Println("Escuchando localhost:8080...")
	// server.ListenAndServe()

	log.Println("Escuchando localhost:9000...")
	err := http.ListenAndServe(":9000", handlers.CORS(header, methods, origins)(r))
	if err != nil {
		fmt.Println(err)
	}
}
