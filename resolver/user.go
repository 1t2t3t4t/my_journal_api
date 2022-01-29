package resolver

import "github.com/1t2t3t4t/my_journal_api/service"

func (r *Resolver) User(arg struct{ Uid string }) (*UserResolver, error) {
	user, err := r.userService.GetUser(arg.Uid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	res := newUserResolver(*user)
	return &res, nil
}

func (r *Resolver) Register(arg struct {
	Username string
	Password string
}) (UserResolver, error) {
	user, err := r.userService.Register(arg.Username, arg.Password)
	if err != nil {
		return UserResolver{}, err
	}
	return newUserResolver(user), nil
}

type UserResolver struct {
	service.User
}

func newUserResolver(user service.User) UserResolver {
	return UserResolver{user}
}
