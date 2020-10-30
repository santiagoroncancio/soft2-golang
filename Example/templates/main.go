package main

import (
	"html/template"
	"os"
)

type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	persona := []Persona{
		{"Santiago", 22},
		{"Pedro", 21},
		{"Alejandro", 23},
	}
	t := template.New("persona")
	t, err := t.ParseGlob("template/*.txt")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(os.Stdout, "visitante", persona)
	if err != nil {
		panic(err)
	}
}
