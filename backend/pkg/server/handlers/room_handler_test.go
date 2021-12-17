package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/novahack/backend/pkg/logging"

	"github.com/matryer/is"
	"github.com/novahack/backend/pkg/database"
	"github.com/novahack/backend/pkg/database/models"
	"github.com/novahack/backend/pkg/database/repos"
	"github.com/novahack/backend/pkg/server"
)

type mockRoomRepo struct {
	createCallCount int
}

func (r *mockRoomRepo) Create() string {
	r.createCallCount++
	return "new-room-uuid"
}

func (r *mockRoomRepo) StartTimer(id string) (int64, error) {
	return 0, nil
}

func (r *mockRoomRepo) Players(id string) ([]models.Player, error) {
	return nil, nil
}

func TestRoomHandlerCreateRoom(t *testing.T) {
	svr := server.New(log.New(), database.NewConn())

	is := is.New(t)

	req := httptest.NewRequest("POST", "/v1/room", nil)
	rec := httptest.NewRecorder()

	svr.ServeHTTP(rec, req)

	is.Equal(rec.Result().StatusCode, http.StatusCreated)
	is.True(len(rec.Body.String()) > 0)
}

func TestRoomHandlerStartTimer(t *testing.T) {
	is := is.New(t)

	dbConn := database.NewConn()

	roomRepo := repos.NewRoomRepo(dbConn)
	id := roomRepo.Create()

	svr := server.New(log.New(), dbConn)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/v1/room/%s/start-game", id), nil)
	rec := httptest.NewRecorder()

	svr.ServeHTTP(rec, req)

	is.Equal(rec.Code, http.StatusOK)
}
