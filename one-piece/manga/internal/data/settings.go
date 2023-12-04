package data

// TODO: move to ../settings/settings.go

import (
	"path/filepath"
	"time"
)

// Settings
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

// NewSettings
func NewSettings() *Settings {
	return &Settings{
		MaxChapterDownloadsPerDay: 10,
		DownloadDelay:             60000 * 30, // NOTE: 30 min
		DataDownloadDir:           filepath.Join("data", "downloads"),
		FetchWeekDay:              time.Thursday,
		FetchHour:                 18,
	}
}
