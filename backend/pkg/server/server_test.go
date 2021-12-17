package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/novahack/backend/pkg/database"
	log "github.com/novahack/backend/pkg/logging"
	"github.com/novahack/backend/pkg/server"
)

func TestNewServerNotNil(t *testing.T) {
	svr := server.New(log.New(), database.NewConn())

	is := is.New(t)
	is.True(svr != nil) // server instance must not be nil
}

func TestNewServerCreateRoomEndpoint(t *testing.T) {
	svr := server.New(log.New(), database.NewConn())

	is := is.New(t)

	req := httptest.NewRequest("POST", "/v1/room", nil)
	rec := httptest.NewRecorder()

	svr.ServeHTTP(rec, req)

	is.Equal(rec.Result().StatusCode, http.StatusCreated)
	is.True(len(rec.Body.String()) > 0)
}
