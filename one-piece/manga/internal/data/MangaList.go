package data

import "time"

// MangaList contains all data for chapters (and arcs).
type MangaList struct {
	Time     time.Time  `json:"time"`
	Chapters []*Chapter `json:"chapters"`
}

// NewMangaList contains all data for chapters (and arcs).
func NewMangaList() *MangaList {
	return &MangaList{
		Time:     time.Now(),
		Chapters: make([]*Chapter, 0),
	}
}
