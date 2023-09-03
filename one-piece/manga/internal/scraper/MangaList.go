package scraper

import (
	"fmt"
	"log"
	"op-manga-dl/internal/data"
	"time"

	"github.com/gocolly/colly/v2"
)

// ParseMangaList
// :d: (optional) use existing data for parsing the manga list
func ParseMangaList(d *data.MangaList) (*data.MangaList, error) {
	if d == nil {
		d = data.NewMangaList()
	}

	d.Time = time.Now()

	c := colly.NewCollector()

	// TODO: css selector to search for
	c.OnHTML("<css selector>", func(e *colly.HTMLElement) {
		// ...
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("[DEBUG] request to \"%s\"\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("[ERROR] %s: %s\n", r.Request.URL, err)
	})

	c.Visit("http://onepiece-tube.com/manga/kapitel-mangaliste")

	return d, fmt.Errorf("Under construction!")
}
