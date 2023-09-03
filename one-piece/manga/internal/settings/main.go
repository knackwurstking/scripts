package settings

import (
	"encoding/json"
	"log"
	"op-manga-dl/internal/data"
	"os"
)

var (
	MaxChapterDownloadsPerDay int
	DownloadDelay             int

	Settings *data.Settings
)

func init() {
	file, err := os.Open(data.SettingsLocation)
	if err != nil {
		log.Printf("[WARN] Failed to load configuration from \"%s\"!", data.SettingsLocation)
		return
	}

	if err := json.NewDecoder(file).Decode(Settings); err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
}
