package models

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
	DOP    string `json:"publication_date"`
}

//func (b Book) PublicationDateStr() string {
//	return b.DOP.Format("2006-01-02")
//}
