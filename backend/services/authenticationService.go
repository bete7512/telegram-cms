package services

import (
	"log"

	"github.com/bete7512/telegram-cms/models"
	"github.com/bete7512/telegram-cms/utils"
)

func (u *UserService) SignUp(user models.SignupRequest) (models.User, error) {
	password, _ := utils.HashPassword(user.Password)

	userModel := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  password,
	}
	newUser, err := u.UserRepository.Create(userModel)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (u *UserService) Login(email string, password string) (accessToken string, errr error) {
	// implement here login logic
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == (models.User{}) {
		return "", utils.ErrUserNotFound
	}

	if !utils.ComparePassword(password, user.Password) {
		return "", utils.ErrWrongPassword
	}

	accessToken, err = utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (u *UserService) ForgetPassword(email string) error {
	// implement here forget password logic
	user, err := u.UserRepository.FindByEmail(email)
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
