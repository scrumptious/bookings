package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/scrumptious/bookings/pkg/config"
	"github.com/scrumptious/bookings/pkg/handlers"
	"github.com/scrumptious/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const port string = ":8080"

var app config.AppConfig

func main() {
	app.InProduction = false

	app.Session = scs.New()
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Persist = true
	app.Session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = !app.InProduction
	render.NewTemplates(&app)

	r := handlers.NewRepo(&app)
	handlers.NewHandlers(r)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
		//Handler:      mux,
		//ReadTimeout:  5,
		//WriteTimeout: 15,
		//IdleTimeout:  15,
	}
	fmt.Println("Starting listening on ", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
