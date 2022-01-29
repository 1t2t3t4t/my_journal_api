package service

import (
	"github.com/1t2t3t4t/my_journal_api/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Uid      string
	Username string
}

type UserService interface {
	Register(username, password string) (User, error)
}

func NewUserService(repository database.UserRepository) UserService {
	return &userService{
		userRepository: repository,
	}
}

const userPasswordHashCost = 12

type userService struct {
	userRepository database.UserRepository
}

func (u *userService) Register(username, password string) (User, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), userPasswordHashCost)
	if err != nil {
		return User{}, err
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return User{}, err
	}

	user, err := u.userRepository.Register(uid.String(), username, string(hashedPwd))
	if err != nil {
		return User{}, err
	}

	return autoCreateMap[User](user)
}
