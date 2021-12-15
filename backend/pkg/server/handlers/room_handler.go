package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/novahack/backend/pkg/database/repos"
	"github.com/novahack/backend/pkg/logging"
)

func CreateRoom(log logging.Logger, repo repos.Room) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		roomID := repo.Create()
		log.Info().Msgf("Created room of ID: %s", roomID)
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(roomID))
	}
}

// PROTOTYPE NOTE: regardless of given room id, player list
// will always be the same but the id of the room does need to exist
func RoomPlayers(log logging.Logger, repo repos.Room) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// distinguish between not found room or other
		// transient or permenant error to affect status code
		players, err := repo.Players(chi.URLParam(r, "uuid"))
		if err != nil {
			rw.Write([]byte(err.Error()))
			return
		}

		writeJSONResp(rw, players)
	}
}

func writeJSONResp(rw http.ResponseWriter, body interface{}) error {
	rw.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(body)
}
