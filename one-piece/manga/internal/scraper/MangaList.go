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

	c.OnHTML("script", func(e *colly.HTMLElement) {
		e.Text = strings.Trim(e.Text, " ")
		if len(e.Text) < 13 {
			return
		}

		if e.Text[0:13] == "window.__data" {
			e.Text, _ = strings.CutPrefix(e.Text, "window.__data")
			e.Text = strings.TrimLeft(e.Text, " =")
			e.Text = strings.TrimRight(e.Text, "; ")

			//os.WriteFile(filepath.Join("data", "window-data.parsed.json"), []byte(t), 0644)

			if err := json.Unmarshal([]byte(e.Text), data); err != nil {
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
