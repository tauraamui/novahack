package handlers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/novahack/backend/pkg/logging"
	"github.com/novahack/backend/pkg/server/handlers"
)

type mockRoomRepo struct {
	createCallCount int
}

func (r *mockRoomRepo) Create() string {
	r.createCallCount++
	return "new-room-uuid"
}

func TestRoomHandlerCreateRoom(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	roomRepo := mockRoomRepo{}
	handlers.CreateRoom(logging.New(), &roomRepo)(rec, req)

	is.Equal(roomRepo.createCallCount, 1)
	is.Equal(rec.Body.String(), "new-room-uuid")
}
