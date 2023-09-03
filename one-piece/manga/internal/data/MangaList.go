package data

type Special struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Arc struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type Chapter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Number      int    `json:"number"`
	CategoryId  int    `json:"category_id"`
	ArcId       int    `json:"arc_id"`
	SpecialsId  int    `json:"specials_id"`
	Lang        string `json:"lang"`
	Pages       int    `json:"pages"`
	IsAvailable bool   `json:"is_available"`
	Date        string `json:"date"`
	Href        string `json:"href"`
}

type MangaList struct {
	Specials []Special `json:"specials"`
	Arcs     []Arc     `json:"arcs"`
	Chapters []Chapter `json:"entries"`
}

func (mangaList *MangaList) GetArc(id int) (arc *Arc, index int) {
	for i := 0; i < len(mangaList.Arcs); i++ {
		if mangaList.Arcs[i].Id == id {
			return &mangaList.Arcs[i], i
		}
	}

	return nil, -1
}
