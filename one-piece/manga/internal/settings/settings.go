package settings

import (
	"encoding/json"
	"log"
	"op-manga-dl/internal/data"
	"os"
	"path/filepath"
	"time"
)

var (
	SettingsLocation string = filepath.Join("data", "settings.json")
	CacheDir         string = "data"

	MaxChapterDownloadsPerDay int
	DownloadDelay             int
	DataDownloadDir           string
	FetchDate                 time.Weekday
	FetchTime                 int
)

func init() {
	file, err := os.Open(SettingsLocation)
	if err != nil {
		log.Printf("[WARN] Failed to load configuration from \"%s\"!", SettingsLocation)
		return
	}

	var s *data.Settings = data.NewSettings()
	if err := json.NewDecoder(file).Decode(s); err != nil {
		log.Fatalf("[FATAL] %s", err)
	}

	MaxChapterDownloadsPerDay = s.MaxChapterDownloadsPerDay
	DownloadDelay = s.DownloadDelay
	DataDownloadDir = s.DataDownloadDir
	FetchDate = s.FetchDate
	FetchTime = s.FetchTime
}
