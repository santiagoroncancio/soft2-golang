package main

import (
	"os"
	"text/template"
)

const HERO = `
Hero Name: {{.Name}}
{{range .Emails}}
Email: {{.}}
{{end}}
{{with .Friends}}
{{range .}}
Friend Name: {{.Name}}
{{end}}
{{end}}
`

type Friend struct {
	Name string
}

type Hero struct {
	Name    string
	Emails  []string
	Friends []Friend
}

func main() {
	f1 := Friend{"santiago"}
	f2 := Friend{"Paco"}
	t := template.New("HERO")
	t, err := t.Parse(HERO)
	if err != nil {
		panic(err)
	}

	hero := Hero{
		Name:    "MoonCake",
		Emails:  []string{"Correo@correo.com", "Correo2@correo2.com"},
		Friends: []Friend{f1, f2},
	}
	err = t.Execute(os.Stdout, hero)
	if err != nil {
		panic(err)
	}
}
