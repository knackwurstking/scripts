package settings

import (
	"encoding/json"
	"log"
	"op-manga-dl/internal/data"
	"os"
	"time"
)

var (
	MaxChapterDownloadsPerDay int
	DownloadDelay             int
	DataDownloadDir           string
	FetchDate                 = time.Thursday
	FetchTime                 = 18
)

func init() {
	file, err := os.Open(data.SettingsLocation)
	if err != nil {
		log.Printf("[WARN] Failed to load configuration from \"%s\"!", data.SettingsLocation)
		return
	}

	var s *data.Settings = data.NewSettings()
	if err := json.NewDecoder(file).Decode(s); err != nil {
		log.Fatalf("[FATAL] %s", err)
	}

	MaxChapterDownloadsPerDay = s.MaxChapterDownloadsPerDay
	DownloadDelay = s.DownloadDelay
	DataDownloadDir = s.DataDownloadDir
}
