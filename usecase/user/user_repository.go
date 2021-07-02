package user

import (
	"github.com/satorunooshie/user-rest-api/domain"
)

type UserRepository interface {
	Find(int) (domain.User, error)
	List() (domain.Users, error)
	Store(domain.User) error
	Update(domain.User) error
	Delete(int) error
}
