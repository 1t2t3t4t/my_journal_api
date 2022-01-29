package resolver

import "github.com/1t2t3t4t/my_journal_api/service"

func (r *Resolver) Login(arg struct {
	Username string
	Password string
}) (UserSessionResolver, error) {
	user, err := r.userService.Login(arg.Username, arg.Password)
	if err != nil {
		return UserSessionResolver{}, err
	}
	return createResult(user)
}

func (r *Resolver) Register(arg struct {
	Username string
	Password string
}) (UserSessionResolver, error) {
	user, err := r.userService.Register(arg.Username, arg.Password)
	if err != nil {
		return UserSessionResolver{}, err
	}
	return createResult(user)
}

func createResult(user service.User) (UserSessionResolver, error) {
	userResolver := newUserResolver(user)
	authToken, err := service.CreateAuthToken(service.AuthClaim{Uid: user.Uid})
	if err != nil {
		return UserSessionResolver{}, err
	}
	return UserSessionResolver{
		userResolver,
		Session{Token: authToken},
	}, nil
}

type Session struct {
	Token string
}

type UserSessionResolver struct {
	UserResolver
	Session Session
}
