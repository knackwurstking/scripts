package scraper

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

type ChapterData_Page struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Type   string `json:"type"` // "image/png" | "image/jpeg"
}

type ChapterData_Chapter struct {
	Name  string             `json:"name"`
	Pages []ChapterData_Page `json:"pages"`
}

type ChapterData struct {
	Chapter ChapterData_Chapter `json:"chapter"`
}

func ParseChapter(chapterURL string) (data *ChapterData, err error) {
	data = &ChapterData{}
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
