package scraper

import (
	"encoding/json"
	"log"
	Data "op-manga-dl/internal/data"
	"strings"

	"github.com/gocolly/colly/v2"
)

// ParseMangaList
func ParseMangaList() (data *Data.MangaList, err error) {
	data = &Data.MangaList{}
	c := colly.NewCollector()

	// TODO: css selector to search for
	// TODO: scrape for arc name: ".segment-heading .segment-name"
	// TODO: scrape for chapters: ".segment-body .segment-row" => ".segment-(number|name|date|pages|lang)"
	c.OnHTML("script", func(e *colly.HTMLElement) {
		if len(e.Text) < 13 {
			return
		}

		if e.Text[0:13] == "window.__data" {
			t := e.Text
			t, _ = strings.CutPrefix(t, "window.__data")
			t = strings.TrimLeft(t, " =")
			t = strings.TrimRight(t, "; ")

			//os.WriteFile(filepath.Join("data", "window-data.parsed.json"), []byte(t), 0644)

			if err := json.Unmarshal([]byte(t), data); err != nil {
				log.Printf("[ERROR] %s\n", err)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("[DEBUG] request to \"%s\"\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		err = e
	})

	if err := c.Visit("https://onepiece-tube.com/manga/kapitel-mangaliste"); err != nil {
		return data, err
	}

	c.Wait()

	return data, err
}
