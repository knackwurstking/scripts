package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"op-manga-dl/internal/scraper" // this will load the configuration (settings.json) file
	"op-manga-dl/internal/settings"
	"op-manga-dl/internal/utils"
	"os"
	"path/filepath"
	"time"
)

func main() {
    downloadAllChapters()

    for true {
        now := time.Now()
        next := time.Date(now.Year(), now.Month(), now.Day()+1, settings.FetchHour, 0, 0, 0, time.Local)

        duration := next.Sub(now)
        log.Printf("sleep a while: %d", duration)
        time.Sleep(duration)

        if next.Weekday() == settings.FetchWeekDay {
            downloadAllChapters()
        }
    }
}

func downloadAllChapters() {
    log.Printf("download all chapters available")

	ml, err := scraper.ParseMangaList()
	if err != nil {
        log.Printf("error: fetch & parse manga list: %s", err)
        return
	}

	for _, chapter := range ml.Chapters {
		arc, i := ml.GetArc(chapter.ArcId)
		if arc == nil {
            log.Fatalf(
                "fatal: Arc for %s with the id %d not found! (This should never happen)",
				chapter.Name, chapter.ArcId,
			)
		}

		path := filepath.Join(
			settings.DataDownloadDir,
			fmt.Sprintf("%03d %s", len(ml.Arcs)-i, arc.Name),
			fmt.Sprintf("%04d %s", chapter.Number, chapter.Name),
		)

		_, err := os.Stat(path + ".pdf")
		if err != nil {
			_ = os.MkdirAll(path, 0755)
			downloadChapter(chapter, path)

			if time.Now().Weekday() == settings.FetchWeekDay &&
				time.Now().Hour() >= settings.FetchHour {

                newMl, err := scraper.ParseMangaList()
				if err != nil {
                    log.Printf("error: %s", err)
				} else {
                    ml = newMl
                }
			}

			sleep()
		}
	}
}

func downloadChapter(chapter scraper.MangaList_Chapter, path string) {
    log.Printf("download the %d pages: \"%s\"", chapter.Pages, path)

	// download jpg/png from dURL - scrape the same script section like before
	chapterData, err := scraper.ParseChapter(chapter.Href)
	if err != nil {
        log.Printf("error: parse chapter: %s\n", err)
		return
	}

	pages := make([]string, len(chapterData.Chapter.Pages))
	for i, page := range chapterData.Chapter.Pages {
        log.Printf("downloading \"%s\"", page.Url)
		r, err := http.Get(page.Url)
		if err != nil {
            log.Printf("error: error while downloading page %d: %s", i+1, err)
			return
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
            log.Printf("error: reading body data for page %d: %s", i+1, err)
			return
		}
		if len(data) == 0 {
            log.Printf("error: no data to read for page %d", i+1)
			return
		}
		e, _ := utils.GetExtension(page.Type)
		p := filepath.Join(path, fmt.Sprintf("%02d.%s", i+1, e))
		err = os.WriteFile(p, data, 0644)
		if err != nil {
            log.Printf("error: write file \"%s\" failed: %s", p, err)
			return
		}
		pages[i] = p
	}

	if err := utils.ConvertImagesToPDF(path, pages...); err != nil {
        log.Printf("error: Convert pages to pdf failed: %s", err)
	}
}

func sleep() {
	log.Printf("wait %d ms...", settings.DownloadDelay)
	time.Sleep(time.Millisecond * time.Duration(settings.DownloadDelay))
}
