package data

type MangaList_Special struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MangaList_Arc struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type MangaList_Chapter struct {
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
	Specials []MangaList_Special `json:"specials"`
	Arcs     []MangaList_Arc     `json:"arcs"`
	Chapters []MangaList_Chapter `json:"entries"`
}

func (mangaList *MangaList) GetArc(id int) (arc *MangaList_Arc, index int) {
	for i := 0; i < len(mangaList.Arcs); i++ {
		if mangaList.Arcs[i].Id == id {
			return &mangaList.Arcs[i], i
		}
	}

	return nil, -1
}
