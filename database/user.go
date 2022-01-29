package database

type User struct {
	Uid            string
	Username       string
	HashedPassword string
}

type UserRepository interface {
	Register(uid, username, hashedPassword string) (User, error)
}
