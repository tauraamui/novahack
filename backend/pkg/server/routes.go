package server

import (
	"net/http"

	"github.com/novahack/backend/pkg/database/repos"
	log "github.com/novahack/backend/pkg/logging"
	"github.com/novahack/backend/pkg/server/handlers"
)

const APIv1 = "v1"

type RouteCollection map[string][]Route

type Route struct {
	endpoint string
	method   string
	handler  http.Handler
}

var versionedRoutes = func(log log.Logger, roomRepo repos.Room) RouteCollection {
	return map[string][]Route{
		APIv1: {
			{
				endpoint: "/room/{guid}/players",
				method:   "GET",
				handler:  handlers.RoomPlayers(log, roomRepo),
			},
		},
	}
}

func (s *Server) mapRoutes(l log.Logger, version string) {
	l.Info().Msg("Mapping routes...")

	for _, v := range versionedRoutes(l, s.roomRepo)[version] {
		l.Info().Msgf("Mapping route %s [%s]", v.method, v.endpoint)
		s.router.Method(v.method, APIv1+"/"+v.endpoint, v.handler)
	}
	// s.router.Method("/", "GET", func(rw http.ResponseWriter, r *http.Request) {})
}

// func enableCors(next handlerFunc) handlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		if r.Method == http.MethodOptions {
// 			return
// 		}
// 		next(w, r)
// 	}
// }
