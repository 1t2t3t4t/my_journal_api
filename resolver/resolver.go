package resolver

import "github.com/1t2t3t4t/my_journal_api/service"

type Resolver struct {
	userService service.UserService
}

type Services struct {
	UserService service.UserService
}

func NewResolver(services *Services) *Resolver {
	return &Resolver{
		userService: services.UserService,
	}
}
