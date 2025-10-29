package main

import (
	"fmt"
	"net/http"
	"time"
)

type application struct {
	port int
	srv  *http.Server
}

func (app *application) serve(handler http.Handler) error {
	app.srv = &http.Server{
		Addr:              fmt.Sprintf(":%d", app.port),
		Handler:           handler,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	return app.srv.ListenAndServe()
}
