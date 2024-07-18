package services

import (
	"log"

	"github.com/bete7512/telegram-cms/models"
)

func (u *UserService) Login(email string, password string) (models.User, error) {
	// implement here login logic
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	//TODO: implement bcrypt compare password
	return user, nil
}


func (u *UserService) SignUp(user models.User) (models.User, error) {
	// implement here signup logic
	// TODO: implement bcrypt password
	newUser, err := u.UserRepository.Create(user)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (u *UserService) ForgetPassword(email string) error {
	// implement here forget password logic
	user , err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	log.Println(user)
	// TODO: implement  prepare token and send email
	return nil
}

func (u *UserService) ResetPassword(token string, password string) error {
	// implement here reset password logic
	// TODO: implement something from coming token then update password
	return nil
}

func (u *UserService) ChangePassword(user models.User, oldPassword string, newPassword string) error {
	// implement here change password logic
	// TODO: implement bcrypt compare password
	return nil
}

func (u *UserService) VerifyEmail(token string) error {
	// implement here verify email logic
	// TODO: implement something here
	return nil
}