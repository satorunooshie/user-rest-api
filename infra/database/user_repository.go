package database

import (
	"github.com/satorunooshie/user-rest-api/domain"
)

type userRepository struct {
	SQLHandler
}

func NewUserRepository(sqlHandler SQLHandler) *userRepository {
	return &userRepository{
		SQLHandler: sqlHandler,
	}
}

func (r *userRepository) Find(id int) (domain.User, error) {
	var user domain.User
	if err := r.SQLHandler.Find(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) List() (domain.Users, error) {
	var users domain.Users
	if err := r.SQLHandler.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) Store(user domain.User) error {
	if err := r.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user domain.User) error {
	if err := r.SQLHandler.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	if err := r.SQLHandler.Delete(id).Error; err != nil {
		return err
	}
	return nil
}
