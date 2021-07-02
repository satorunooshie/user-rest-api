package user

import (
	"github.com/satorunooshie/user-rest-api/domain"
)

type UserInteractor interface {
	Find(int) (domain.User, error)
	List() (domain.Users, error)
	Store(domain.User) error
	Update(domain.User) error
	Delete(int) error
}

type userInteractor struct {
	repo UserRepository
}

var _ UserInteractor = (*userInteractor)(nil)

func NewUserInteractor(repo UserRepository) UserInteractor {
	return &userInteractor{
		repo: repo,
	}
}

func (i *userInteractor) Find(id int) (domain.User, error) {
	return i.repo.Find(id)
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

func (i *userInteractor) Delete(id int) error {
	return i.repo.Delete(id)
}
