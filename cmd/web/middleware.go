package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(fn)
}

func CsrfMid(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
