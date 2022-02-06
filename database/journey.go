package database

type Journey struct {
	AuthorUid string
	Title     string
	Content   string
}

type JourneyRepository interface {
	Create(authorUid, title, content string) (Journey, error)
}
