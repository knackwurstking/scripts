package data

import "path/filepath"

var (
	SettingsLocation string = filepath.Join("data", "settings.json")
)

// Settings
type Settings struct {
	// MaxChapterDownloadsDerDay to download
	MaxChapterDownloadsDerDay int `json:"max-chapter-downloads-per-day"`
	// DownloadDelay is the time in ms to wait between downloads
	DownloadDelay int `json:"download-delay"`
}
