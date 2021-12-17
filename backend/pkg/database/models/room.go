package models

type Room struct {
	UUID           string
	Players        []Player
	TimerStartTime int64
}

func (r Room) ID() string {
	return r.UUID
}
