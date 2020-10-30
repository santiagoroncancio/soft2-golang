package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func prueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo desde prueba</h1>")
}

func usuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola usuario</h1>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {
	msg := mensaje{
		msg: "<h1>Holaaaaa C:</h1>",
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	mux.HandleFunc("/prueba", prueba)
	mux.HandleFunc("/usuario", usuario)
	mux.Handle("/hola", msg)

	// ServerMux o enrutador
	// http.ListenAndServe(":8080", mux)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Escuchando...")
	log.Fatal(server.ListenAndServe())
}

// http.Handle() <-- Funciones del metodo http
// http.HandleFunc()

// http.Handler <-- interfaz del controlador
// http.HandlerFunc
