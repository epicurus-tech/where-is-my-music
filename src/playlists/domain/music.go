package domain

import "time"

type Music struct {
	Name       string
	Date       time.Time
	Compositor []Compositor
}

func NewMusic(name string, date time.Time, compositor []Compositor) *Music {
	return &Music{
		Name:       name,
		Date:       date,
		Compositor: compositor,
	}
}
