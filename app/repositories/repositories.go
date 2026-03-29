package repositories

import (
	"github.com/create-go-app/fiber-go-template/platform/database"
)

type Repositories struct {
	User UserRepository
	Ssh  SshRepository
}

func NewRepositories() Repositories {
	db, _ := database.OpenDBConnection()

	return Repositories{
		User: NewUserRepository(db.UserQueries),
		Ssh:  NewSshRepository(db.SSHQueries),
	}
}
