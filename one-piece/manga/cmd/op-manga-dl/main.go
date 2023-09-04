package main

import (
	"encoding/json"
	"fmt"
	"log"
	Data "op-manga-dl/internal/data"
	"op-manga-dl/internal/scraper" // this will load the configuration (settings.json) file
	"op-manga-dl/internal/settings"
	"os"
	"path/filepath"
	"time"
)

// TODO: ...
// [x] parse manga list for chapters and arcs (cache in data/data.json)
// [x] check for missing chapters in data/downloads
// [x] get next chapter to download
// [-] parse url to chapter and get all available pages
// [ ] download each page (jpg 01-??)
// [ ] merge all jpg's to a pdf with ImageMagic (`convert "*.{jpg}" -quality 100 -density 150 "<nr.> <chapter name>.pdf"`)
// [ ] mark chapter as complete in "data/data.json"

func main() {
	mangaList, err := scraper.ParseMangaList()
	if err != nil {
		log.Fatalf("[FATAL] %s\n", err)
	}

	d, _ := json.MarshalIndent(mangaList, "", "    ")
	os.WriteFile(filepath.Join("data", "data.json"), d, 0644)

	for _, chapter := range mangaList.Chapters {
		arc, i := mangaList.GetArc(chapter.ArcId)
		if arc == nil {
			log.Fatalf(
				"[FATAL] Arc for %s with the id %d not found! (This should never happen)\n",
				chapter.Name, chapter.ArcId,
			)
		}

		path := filepath.Join(
			settings.DataDownloadDir,
			fmt.Sprintf("%03d %s", i, arc.Name),
			fmt.Sprintf("%04d %s", chapter.Number, chapter.Name),
		)

		_, err := os.Stat(path + ".pdf")
		if err != nil {
			// file does not exists, mark download
			download(chapter, path)
			sleep()
		}
	}
}

func download(chapter Data.MangaList_Chapter, path string) {
	log.Printf("[DEBUG] @TODO: download the %d pages to \"%s\"", chapter.Pages, path)

	// download jpg/png from dURL - scrape the same script section like before
	chapterData, err := scraper.ParseChapter()
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
		return
	}

	// TODO: store json data to "<path>/data.json" first

	for i, page := range chapterData.Chapter.Pages {
		// TODO: download image from `page.Url` and save as ("%02d", i)
	}

	log.Printf("[DEBUG] @TODO: and convert to \"%s\"", path+".pdf")
}

func sleep() {
	time.Sleep(time.Millisecond * time.Duration(settings.DownloadDelay))
}
