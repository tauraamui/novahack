package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/novahack/backend/pkg/database"
	"github.com/novahack/backend/pkg/database/repos"
	log "github.com/novahack/backend/pkg/logging"
)

type Server struct {
	router   chi.Router
	roomRepo repos.Room
}

func New(l log.Logger, db database.DBConn) *Server {
	svr := &Server{
		router:   chi.NewRouter(),
		roomRepo: repos.NewRoomRepo(db),
	}
	svr.mapRoutes(l, APIv1)
	return svr
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
