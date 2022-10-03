package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marceesty/learning_go_web/pkg/config"
	"github.com/marceesty/learning_go_web/pkg/handlers"
	"github.com/marceesty/learning_go_web/pkg/render"
)

const portNumber = ":8080"


func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting app on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}