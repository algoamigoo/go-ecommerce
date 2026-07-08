package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// mount the API
func (app *application) mount() http.Handler{
	r := chi.NewRouter()

	// user -> handler Get / products -> service getProducts -> throw errors

	// a good base middleware stack
	r.Use(middleware.RequestID)		//important for rate limiting
	r.Use(middleware.RealIP)	// used for rate limiting and analytics and tracing
	r.Use(middleware.Logger)	// 
	r.Use(middleware.Recoverer)		// recover from crashes
	
	// sets timeout for request context (ctx) and will through ctx.Done() that the request has timed out and further processing should be stopped
	r.Use(middleware.Timeout(60*time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	// http.ListenAndServe(":3333", r)

	return r
}

// run 

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	log.Printf("Server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}


type application struct { 
	config config
	// logger
	// db driver

}

type config struct {
	addr string // port no
	db dbConfig
}

type dbConfig struct {
	dsn string 
}