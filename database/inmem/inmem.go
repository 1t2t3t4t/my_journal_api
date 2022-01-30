package inmem

import "github.com/1t2t3t4t/my_journal_api/database"

func DefaultRepositories() *database.Repositories {
	return &database.Repositories{
		UserRepository: newUserRepositories(),
	}
}
