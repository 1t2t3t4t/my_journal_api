package database

import "github.com/1t2t3t4t/my_journal_api/types"

type Journey struct {
	AuthorUid string
	Title     string
	Content   string
	CreatedAt types.DateTime
	UpdatedAt types.DateTime
}

type JourneyRepository interface {
	Create(authorUid, title, content string) (Journey, error)
	FindAll(authorUid string) ([]Journey, error)
}
