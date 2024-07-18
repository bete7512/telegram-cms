package services

import (
	"github.com/bete7512/telegram-cms/models"
	"github.com/bete7512/telegram-cms/ports"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) FindAll() ([]models.User, error) {
	return u.UserRepository.FindAll()
}

func (u *UserService) FindByID(id int) (models.User, error) {
	return u.UserRepository.FindByID(id)
}

func (u *UserService) Create(user models.User) (models.User, error) {
	return u.UserRepository.Create(user)
}

func (u *UserService) Update(user models.User) (models.User, error) {
	return u.UserRepository.Update(user)
}

func (u *UserService) Delete(id int) (bool, error) {
	return u.UserRepository.Delete(id)
}

