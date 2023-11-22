package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const Version = "1"
const Port = "8080"

type HomePageVars struct {
	Version string
}

type Application struct {
	templates *template.Template
}

func (app *Application) HomePage(w http.ResponseWriter, r *http.Request) {
	pageVars := HomePageVars{
		Version: Version,
	}
	app.templates.ExecuteTemplate(w, "homepage", pageVars)
}

func main() {
	app := Application{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	http.HandleFunc("/", app.HomePage)
	err := http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
	if err != nil {
		panic(err)
	}
}
