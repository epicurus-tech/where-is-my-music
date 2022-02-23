package domain

import "time"

type Album struct {
	Name     string
	Duration time.Duration
	Musics   []Music
}

func NewAlbum(name string, duration time.Duration, musics []Music) *Album {
	return &Album{
		Name:     name,
		Duration: duration,
		Musics:   musics,
	}
}
