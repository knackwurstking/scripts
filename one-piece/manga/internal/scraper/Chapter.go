package scraper

import (
	"encoding/json"
	"log"
	Data "op-manga-dl/internal/data"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ParseChapter(chapterURL string) (data *Data.ChapterData, err error) {
	data = &Data.ChapterData{}
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

	if err := c.Visit(chapterURL); err != nil {
		return data, err
	}

	c.Wait()

	return data, err
}
