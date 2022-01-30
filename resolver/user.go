package resolver

import (
	"context"

	"github.com/1t2t3t4t/my_journal_api/service"
)

func (r *Resolver) User(ctx context.Context, arg struct{ Uid *string }) (*UserResolver, error) {
	var targetUid string
	if arg.Uid != nil {
		targetUid = *arg.Uid
	} else if claim, ok := guardLoggedInUser(ctx); ok {
		targetUid = claim.Uid
	} else {
		return nil, ResolvingErrorNotLoggedInUser
	}

	user, err := r.userService.GetUser(targetUid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	res := newUserResolver(*user)
	return &res, nil
}

type UserResolver struct {
	service.User
}

func newUserResolver(user service.User) UserResolver {
	return UserResolver{user}
}
