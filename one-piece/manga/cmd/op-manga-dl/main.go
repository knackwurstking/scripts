package main

import (
	"encoding/json"
	"fmt"
	"log"
	"op-manga-dl/internal/scraper" // this will load the configuration (settings.json) file
	"op-manga-dl/internal/settings"
	"os"
	"path/filepath"
)

// TODO: ...
// [x] parse manga list for chapters and arcs (cache in data/data.json)
// [ ] check for missing chapters in data/downloads
// [ ] get next chapter to download
// [ ] parse url to chapter and get all available pages
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

		// TODO: ...
		log.Printf("[DEBUG] @TODO: check if chapter \"%s\" exists", path)
	}
}
