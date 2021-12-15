package handlers

import (
	"net/http"

	"github.com/novahack/backend/pkg/database/repos"
	"github.com/novahack/backend/pkg/logging"
)

func CreateRoom(log logging.Logger, repo repos.Room) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		roomID := repo.Create()
		log.Info().Msgf("Created room of ID: %s", roomID)
		rw.Write([]byte(roomID))
	}
}

func RoomPlayers(log logging.Logger, repo repos.Room) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}
