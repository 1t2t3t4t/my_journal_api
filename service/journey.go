package service

import (
	"github.com/1t2t3t4t/my_journal_api/database"
	"github.com/1t2t3t4t/my_journal_api/types"
)

type Journey struct {
	Title     string
	Content   string
	CreatedAt types.DateTime
	UpdatedAt types.DateTime
}

type JourneyService interface {
	Create(authorUid, title, content string) (Journey, error)
	FindAll(authorUid string) ([]Journey, error)
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

func (j *journeyService) FindAll(authorUid string) ([]Journey, error) {
	journeys, err := j.repository.FindAll(authorUid)
	if err != nil {
		return nil, err
	}
	return Map(journeys, func(v database.Journey) (Journey, bool) {
		res, err := autoCreateMap[Journey](v)
		if err != nil {
			return Journey{}, false
		}
		return res, true
	}), nil
}
