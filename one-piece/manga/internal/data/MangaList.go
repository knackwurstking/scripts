package data

// MangaList contains all data for chapters (and arcs).
type MangaList struct {
	Chapters []*Chapter `json:"chapters"`
}
