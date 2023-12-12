package anime

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	PathEpisodenStreams Path = "/anime/episoden-streams"
)

type Path string

type Anime struct {
	Origin string `json:"origin"`
	Data   *Data  `json:"data"`
}

func New(origin string) *Anime {
	return &Anime{
		Origin: origin,
		Data:   NewData(),
	}
}

func (anime *Anime) GetUrl(path Path) string {
	switch path {
	case PathEpisodenStreams:
		return fmt.Sprintf("%s%s", anime.Origin, PathEpisodenStreams)
	default:
		panic(fmt.Sprintf("Name \"%s\" not found!", path))
	}
}

func (anime *Anime) GetEpisodenStreams() (*Data, error) {
	var (
		c   = colly.NewCollector()
		err error
	)

	c.OnHTML("script", func(h *colly.HTMLElement) {
		var (
			dataVar = "window.__data"
			text    = strings.Trim(h.Text, " ")
		)

		if len(text) < len(dataVar) {
			return
		}

		if text[0:13] == dataVar {
			text, _ = strings.CutPrefix(text, dataVar)
			text = strings.TrimLeft(text, " =")
			text = strings.TrimRight(text, "; ")

			if err := json.Unmarshal([]byte(text), anime.Data); err != nil {
				slog.Error("Unmarshal data failed!", "err", err.Error())
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		slog.Debug(fmt.Sprintf("Request to \"%s\"", r.URL))
	})

	c.OnError(func(r *colly.Response, e error) {
		if e != nil {
			err = e
		}
	})

	if err = c.Visit(anime.GetUrl(PathEpisodenStreams)); err != nil {
		return anime.Data, err
	}

	c.Wait()

	return anime.Data, err
}

type Data struct {
	Category DataCategory `json:"category"`
	Arcs     DataArcs     `json:"arcs"`
	Entries  []DataEntry  `json:"entries"`
}

func NewData() *Data {
	return &Data{
		Category: DataCategory{},
		Arcs:     make([]DataArc, 0),
		Entries:  make([]DataEntry, 0),
	}
}

type DataCategory struct {
	ID   int    `json:"id"`
	Type string `json:"type"` // NOTE: "anime"
}

type DataArcs []DataArc

func (arcs DataArcs) Get(id int) *DataArc {
	for i := 0; i < len(arcs); i++ {
		if arcs[i].ID == id {
			return &arcs[i]
		}
	}

	return nil
}

// GetIndex in reversed order
func (arcs DataArcs) GetIndex(id int) int {
	for i := 0; i < len(arcs); i++ {
		if arcs[i].ID == id {
			return len(arcs) - 1 - i
		}
	}

    panic(fmt.Sprintf("GetOrder failed for id \"%d\"", id))
}

type DataArc struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DataEntry struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	CategoryID  int    `json:"category_id"` // NOTE: only cateogory 1 "type: anime"
	ArcID       int    `json:"arc_id"`
	IsSpecial   bool   `json:"is_special"`
	IsFiller    bool   `json:"is_filler"`
	LangSub     string `json:"lang_sub"`
	LangDub     string `json:"lang_dub"`
	IsAvailable bool   `json:"is_available"`
	Href        string `json:"href"`
}
