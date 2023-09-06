package data

import "path/filepath"

var (
	// SettingsLocation
	SettingsLocation string = filepath.Join("data", "settings.json")
)

// Settings
type Settings struct {
	// MaxChapterDownloadsPerDay to download
	MaxChapterDownloadsPerDay int `json:"max-chapter-downloads-per-day"`
	// DownloadDelay is the time in ms to wait between downloads
	DownloadDelay int `json:"download-delay"`
	// DataDir where to store all download data
	DataDownloadDir string `json:"data-download-dir"`
}

// NewSettings
func NewSettings() *Settings {
	return &Settings{
		MaxChapterDownloadsPerDay: 10,
		DownloadDelay:             60000 * 30, // NOTE: 30 min
		DataDownloadDir:           filepath.Join("data", "downloads"),
	}
}
