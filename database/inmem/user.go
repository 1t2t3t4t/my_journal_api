package inmem

import (
	"github.com/1t2t3t4t/my_journal_api/database"
	"github.com/1t2t3t4t/my_journal_api/utils"
)

type userKey string

type userRepository struct {
	users map[userKey]database.User // Map pair for user's uid and user
}

func NewUserRepository() database.UserRepository {
	return &userRepository{users: make(map[userKey]database.User)}
}

func (r *userRepository) Register(uid, username, hashedPassword string) (database.User, error) {
	if utils.ContainKey(r.users, userKey(uid)) {
		return database.User{}, database.ErrorDuplicateUsername
	}
	user := database.User{Uid: uid, Username: username, HashedPassword: hashedPassword}
	r.users[userKey(uid)] = user
	return user, nil
}

func (r *userRepository) FindOne(uid string) *database.User {
	if user, ok := r.users[userKey(uid)]; ok {
		return &user
	}

	return nil
}
