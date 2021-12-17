package repos

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/novahack/backend/pkg/database"
	"github.com/novahack/backend/pkg/database/models"
)

type Room interface {
	Create() string
	StartTimer(string) (int64, error)
	Players(string) ([]models.Player, error)
}

type roomRepo struct {
	idKeyPrefix string
	db          database.DBConn
}

func NewRoomRepo(db database.DBConn) Room {
	return &roomRepo{
		"room:",
		db,
	}
}

func (r *roomRepo) Create() string {
	return r.db.Create(r.idKeyPrefix, models.Room{
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

func (r *roomRepo) StartTimer(id string) (int64, error) {
	whenTimerWillStart := time.Now().Add(time.Second * 10).Unix()
	if room, ok := r.db.Find(id).(models.Room); ok {
		room.TimerStartTime = whenTimerWillStart
		r.db.Replace(room)
	} else {
		return 0, fmt.Errorf("unable to find room of ID: %s", id)
	}
	return whenTimerWillStart, nil
}

func (r *roomRepo) Players(id string) ([]models.Player, error) {
	// cast from interface to room model with safety check
	if room, ok := r.db.Find(id).(models.Room); ok {
		return room.Players, nil
	}

	return nil, fmt.Errorf("unable to find room %s", id)
}
