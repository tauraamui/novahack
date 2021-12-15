package models

type Player struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Stake int    `json:"stake"`
}
