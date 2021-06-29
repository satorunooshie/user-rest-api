package database

import (
	"github.com/satorunooshie/user-rest-api/domain"
)

type userRepository struct {
	SQLHandler
}

func (r *userRepository) FindByID(id int) (domain.User, error) {
	var user domain.User
	if err := r.Find(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) List() (domain.Users, error) {
	var users domain.Users
	if err := r.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) Store(user domain.User) (domain.User, error) {
	if err := r.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user domain.User) error {
	if err := r.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteByID(user domain.User) error {
	if err := r.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
