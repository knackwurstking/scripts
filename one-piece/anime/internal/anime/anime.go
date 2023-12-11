package anime

import "fmt"

const (
    NameAnimeList Name = "anime-list"
)

type Chapter struct {
}

type Name string

type Anime struct {
	Origin string
}

func New(origin string) *Anime {
	return &Anime{
		Origin: origin,
	}
}

func (anime *Anime) GetUrl(name Name) string {
	switch name {
    case "anime-list":
        return fmt.Sprintf("%s/anime-list", anime.Origin) // TODO: check url
	default:
		panic(fmt.Sprintf("Name \"%s\" not found!", name))
	}
}

func (anime *Anime) GetAnimeList() ([]Chapter, error) {
	// TODO: fetch anime list

	return make([]Chapter, 0), nil
}
