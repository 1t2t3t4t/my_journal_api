package service

import "github.com/1t2t3t4t/my_journal_api/database"

type Journey struct {
	Title   string
	Content string
}

type JourneyService interface {
	Create(authorUid, title, content string) (Journey, error)
}

func NewJourneyService(repository database.JourneyRepository) JourneyService {
	return &journeyService{repository: repository}
}

type journeyService struct {
	repository database.JourneyRepository
}

func (j *journeyService) Create(authorUid, title, content string) (Journey, error) {
	journey, err := j.repository.Create(authorUid, title, content)
	if err != nil {
		return Journey{}, err
	}
	return autoCreateMap[Journey](journey)
}
