package inmem

import "github.com/1t2t3t4t/my_journal_api/database"

type journeyRepository struct {
	journeys map[string][]database.Journey
}

func newJourneyRepository() database.JourneyRepository {
	return &journeyRepository{
		journeys: make(map[string][]database.Journey),
	}
}

func (j *journeyRepository) Create(authorUid, title, content string) (database.Journey, error) {
	journeys := j.safeGetJourneys(authorUid)
	journey := database.Journey{
		AuthorUid: authorUid,
		Title:     title,
		Content:   content,
	}
	journeys = append(journeys, journey)
	j.journeys[authorUid] = journeys
	return journey, nil
}

func (j *journeyRepository) safeGetJourneys(uid string) []database.Journey {
	if journeys, ok := j.journeys[uid]; ok {
		return journeys
	}
	journeys := make([]database.Journey, 1)
	j.journeys[uid] = journeys
	return journeys
}
