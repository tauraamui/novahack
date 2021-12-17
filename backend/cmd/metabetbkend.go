package main

import (
	"net/http"
	"time"

	"github.com/novahack/backend/pkg/database"
	log "github.com/novahack/backend/pkg/logging"
	"github.com/novahack/backend/pkg/server"
)

func main() {
	// init logging
	log := log.New()

	server := server.New(log, database.NewConn())
	srv := &http.Server{
		Addr:         "8282",
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 120,
		Handler:      server,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}
