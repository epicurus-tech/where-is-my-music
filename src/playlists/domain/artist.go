package domain

type Artist struct {
	Name   string
	Genre  string
	Musics []Music
	Albums []Album
	Locale string
}

func NewArtist(name string, genre string, musics []Music, albums []Album, locale string) *Artist {
	return &Artist{
		Name:   name,
		Genre:  genre,
		Musics: musics,
		Albums: albums,
		Locale: locale,
	}
}
