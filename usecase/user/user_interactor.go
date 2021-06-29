package user

import (
	"github.com/satorunooshie/user-rest-api/domain"
)

type userInteractor struct {
	repo UserRepository
}

func (i *userInteractor) FindByID(id int) (domain.User, error) {
	return i.repo.FindByID(id)
}

func (i *userInteractor) List() (domain.Users, error) {
	return i.repo.List()
}

func (i *userInteractor) Store(user domain.User) error {
	return i.repo.Store(user)
}

func (i *userInteractor) Update(user domain.User) error {
	return i.repo.Update(user)
}

func (i *userInteractor) DeleteByID(id int) error {
	return i.repo.DeleteByID(id)
}
