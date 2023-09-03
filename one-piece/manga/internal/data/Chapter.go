package data

// Chapter contains the current arc and chapter name and the chapter number.
type Chapter struct {
	Arc  string `json:"arc"`
	Nr   int    `json:"nr"`
	Name string `json:"name"`
}
