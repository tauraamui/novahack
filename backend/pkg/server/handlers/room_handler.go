package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/novahack/backend/pkg/database/models"
	"github.com/novahack/backend/pkg/database/repos"
	"github.com/novahack/backend/pkg/logging"
)

func CreateRoom(log logging.Logger, repo repos.Room) http.HandlerFunc {
	type data struct {
		ID      string          `json:"id"`
		You     string          `json:"your-player-id"`
		Players []models.Player `json:"players"`
	}
	type respBody struct {
		Data  data
		Error string `json:"error"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		roomID := repo.Create()
		log.Info().Msgf("Created room of ID: %s", roomID)
		rw.WriteHeader(http.StatusCreated)

		roomPlayers, err := repo.Players(roomID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			writeJSONResp(rw, respBody{
				data{},
				err.Error(),
			})
			return
		}

		writeJSONResp(rw, respBody{
			data{
				ID:      roomID,
				You:     roomPlayers[0].UUID,
				Players: roomPlayers,
			},
			"",
		})
	}
}

func StartTimer(log logging.Logger, repo repos.Room) http.HandlerFunc {
	type data struct {
		Timestamp int64 `json:"timer-start-time"`
	}

	type respBody struct {
		Data  data
		Error string `json:"error"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		roomID := chi.URLParam(r, "uuid")
		timerStartTime, err := repo.StartTimer(roomID)
		if err != nil {
			// could be 404 or 500 really, but eh
			rw.WriteHeader(http.StatusInternalServerError)
			writeJSONResp(rw, respBody{
				data{},
				err.Error(),
			})
			return
		}

		writeJSONResp(rw, respBody{
			data{
				Timestamp: timerStartTime,
			},
			"",
		})
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
