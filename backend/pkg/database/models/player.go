package models

type Player struct {
	UUID  string `json:"id"`
	Name  string `json:"name"`
	Stake int    `json:"stake"`
}
