package repos

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/novahack/backend/pkg/database"
	"github.com/novahack/backend/pkg/database/models"
)

const idKeyPrefix = "room:"

type Room interface {
	Create() string
	Players(string) ([]models.Player, error)
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
	return r.db.Create(idKeyPrefix, models.Room{
		Players: []models.Player{
			{
				UUID:  uuid.NewString(),
				Name:  "First Player",
				Stake: 350,
			},
			{
				UUID:  uuid.NewString(),
				Name:  "Second Player",
				Stake: 400,
			},
			{
				UUID:  uuid.NewString(),
				Name:  "Third Player",
				Stake: 50,
			},
			{
				UUID:  uuid.NewString(),
				Name:  "Fourth Player",
				Stake: 900,
			},
		},
	})
}

func (r *roomRepo) Players(id string) ([]models.Player, error) {
	// cast from interface to room model with safety check
	if room, ok := r.db.Find(id).(models.Room); ok {
		return room.Players, nil
	}

	return nil, fmt.Errorf("unable to find room %s", id)
}
