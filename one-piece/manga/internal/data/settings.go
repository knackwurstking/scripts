package data

import "path/filepath"

var (
	SettingsLocation string = filepath.Join("data", "settings.json")
)

// Settings
type Settings struct {
	// ChaptersPerDay to download
	ChaptersPerDay int `json:"ChaptersPerDay"`
	// Delay is the time in ms to wait between downloads
	Delay int `json:"delay"`
}
