package data

type ChapterData_Page struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Type   string `json:"type"` // "image/png" | "image/jpeg"
}

type ChapterData_Chapter struct {
	Name  string             `json:"name"`
	Pages []ChapterData_Page `json:"pages"`
}

type ChapterData struct {
	Chapter ChapterData_Chapter `json:"chapter"`
}
