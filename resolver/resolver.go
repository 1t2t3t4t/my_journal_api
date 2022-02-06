package resolver

import "github.com/1t2t3t4t/my_journal_api/service"

type Resolver struct {
	*Services
}

type Services struct {
	UserService    service.UserService
	JourneyService service.JourneyService
}

func NewResolver(services *Services) *Resolver {
	return &Resolver{
		services,
	}
}
