package database

type User struct {
	Uid            string
	Username       string
	HashedPassword string
}

type UserRepository interface {
	FindOne(uid string) *User
	FindOneByUsername(username string) *User

	Register(uid, username, hashedPassword string) (User, error)
}
