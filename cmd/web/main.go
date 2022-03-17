package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brianweber2/golang_web_app/pkg/config"
	"github.com/brianweber2/golang_web_app/pkg/handlers"
	"github.com/brianweber2/golang_web_app/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
