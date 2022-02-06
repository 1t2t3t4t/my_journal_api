package resolver

import (
	"context"

	"github.com/1t2t3t4t/my_journal_api/service"
)

func (r *Resolver) Journeys(ctx context.Context) ([]JourneyResolver, error) {
	userClaim, ok := guardLoggedInUser(ctx)
	if !ok {
		return nil, ResolvingErrorNotLoggedInUser
	}
	journeys, err := r.JourneyService.FindAll(userClaim.Uid)
	if err != nil {
		return nil, err
	}
	user, err := r.UserService.GetUser(userClaim.Uid)
	if err != nil {
		return nil, err
	}
	res := service.Map(journeys, func(journey service.Journey) (JourneyResolver, bool) {
		return NewJourneyResolver(journey, *user), true
	})
	return res, nil
}

func (r *Resolver) CreateJourney(ctx context.Context, args struct {
	Title   string
	Content string
}) (JourneyResolver, error) {
	userClaim, ok := guardLoggedInUser(ctx)
	if !ok {
		return JourneyResolver{}, ResolvingErrorNotLoggedInUser
	}
	j, err := r.JourneyService.Create(userClaim.Uid, args.Title, args.Content)
	if err != nil {
		return JourneyResolver{}, err
	}
	user, err := r.UserService.GetUser(userClaim.Uid)
	if err != nil {
		return JourneyResolver{}, err
	}
	if user == nil {
		return JourneyResolver{}, ResolvingErrorUserNotFound
	}
	return NewJourneyResolver(j, *user), nil
}

func NewJourneyResolver(journey service.Journey, author service.User) JourneyResolver {
	return JourneyResolver{
		journey,
		newUserResolver(author),
	}
}

type JourneyResolver struct {
	service.Journey
	Author UserResolver
}
