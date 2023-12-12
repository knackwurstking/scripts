package anime

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	NameEpisodenStreams Name = "episoden-streams"
)

type Name string

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

func (anime *Anime) GetUrl(name Name) string {
	switch name {
	case NameEpisodenStreams:
		return fmt.Sprintf("%s/anime/episoden-streams", anime.Origin) // TODO: check url
	default:
		panic(fmt.Sprintf("Name \"%s\" not found!", name))
	}
}

func (anime *Anime) GetEpisodenStreams() (*Data, error) {
	anime.Data = nil // reset data first

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

	if err = c.Visit(anime.GetUrl(NameEpisodenStreams)); err != nil {
		return anime.Data, err
	}

	c.Wait()

	return anime.Data, err
}

type Data struct {
	Category DataCategory `json:"category"`
	Arcs     []DataArc    `json:"arcs"`
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

type DataArc struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
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
