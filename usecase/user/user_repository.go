package user

import (
	"github.com/satorunooshie/user-rest-api/domain"
)

type UserRepository interface {
	FindByID(int) (domain.User, error)
	List() (domain.Users, error)
	Store(domain.User) error
	Update(domain.User) error
	DeleteByID(int) error
}
