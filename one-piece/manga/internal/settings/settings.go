package settings

import (
	"encoding/json"
	"log"
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
	FetchWeekDay              time.Weekday
	FetchHour                 int
)

func init() {
	file, err := os.Open(SettingsLocation)
	if err != nil {
		log.Printf("[WARN] Failed to load configuration from \"%s\"!", SettingsLocation)
		return
	}

	var s *Settings = NewSettings()
	if err := json.NewDecoder(file).Decode(s); err != nil {
		log.Fatalf("[FATAL] %s", err)
	}

	MaxChapterDownloadsPerDay = s.MaxChapterDownloadsPerDay
	DownloadDelay = s.DownloadDelay
	DataDownloadDir = s.DataDownloadDir
	FetchWeekDay = s.FetchWeekDay
	FetchHour = s.FetchHour
}

type Settings struct {
	// MaxChapterDownloadsPerDay to download
	MaxChapterDownloadsPerDay int `json:"max-chapter-downloads-per-day"`
	// DownloadDelay is the time in ms to wait between downloads
	DownloadDelay int `json:"download-delay"`
	// DataDir where to store all download data
	DataDownloadDir string `json:"data-download-dir"`
	// FetchWeekDay in range from 0-6
	FetchWeekDay time.Weekday `json:"fetch-date"`
	// FetchHour in range from 0-23
	FetchHour int `json:"fetch-time"`
}

func NewSettings() *Settings {
	return &Settings{
		MaxChapterDownloadsPerDay: 10,
		DownloadDelay:             60000 * 30, // NOTE: 30 min
		DataDownloadDir:           filepath.Join("data", "downloads"),
		FetchWeekDay:              time.Thursday,
		FetchHour:                 18,
	}
}
