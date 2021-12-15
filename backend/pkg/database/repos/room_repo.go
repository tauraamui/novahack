package repos

import (
	"fmt"

	"github.com/novahack/backend/pkg/database"
	"github.com/novahack/backend/pkg/database/models"
)

const idKeyPrefix = "room:"

type Room interface {
	Create() string
}

type roomRepo struct {
	db database.DBConn
}

func NewRoomRepo(db database.DBConn) Room {
	return &roomRepo{
		db,
	}
}

func (r *roomRepo) Create() string {
	return r.db.Create(idKeyPrefix, models.Room{})
}

func (r *roomRepo) Players(id string) ([]models.Player, error) {
	// cast from interface to room model with safety check
	if room, ok := r.db.Find(id).(models.Room); ok {
		return room.Players, nil
	}

	return nil, fmt.Errorf("unable to find room %s", id)
}
