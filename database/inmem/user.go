package inmem

import (
	"sync"

	"github.com/1t2t3t4t/my_journal_api/database"
	"github.com/1t2t3t4t/my_journal_api/utils"
)

type userKey string

type userRepository struct {
	users map[userKey]database.User // Map pair for user's uid and user
}

var userRepositorySingleton *userRepository
var singletonOnce sync.Once
var userLock sync.RWMutex

func NewUserRepository() database.UserRepository {
	singletonOnce.Do(func() {
		userRepositorySingleton = &userRepository{users: make(map[userKey]database.User)}
	})
	return userRepositorySingleton
}

func (r *userRepository) Register(uid, username, hashedPassword string) (database.User, error) {
	userLock.Lock()
	defer userLock.Unlock()
	if utils.ContainKey(r.users, userKey(uid)) {
		return database.User{}, database.ErrorDuplicateUsername
	}
	user := database.User{Uid: uid, Username: username, HashedPassword: hashedPassword}
	r.users[userKey(uid)] = user
	return user, nil
}

func (r *userRepository) FindOne(uid string) *database.User {
	userLock.RLock()
	defer userLock.RUnlock()
	if user, ok := r.users[userKey(uid)]; ok {
		return &user
	}

	return nil
}

func (r *userRepository) FindOneByUsername(username string) *database.User {
	userLock.RLock()
	defer userLock.RUnlock()
	if _, user, found := utils.FindMapValue(r.users, func(v database.User) bool {
		return v.Username == username
	}); found {
		return &user
	}
	return nil
}
