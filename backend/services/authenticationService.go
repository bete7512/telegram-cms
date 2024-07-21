package services

import (
	"fmt"

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
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == (models.User{}) {
		return "", utils.ErrUserNotFound
	}
	if !user.Status {
		return "", utils.ErrUserNotActive
	}
	if !utils.ComparePassword(user.Password, password) {
		return "", utils.ErrWrongPassword
	}

	accessToken, err = utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (u *UserService) ForgetPassword(email string, resetLink string) error {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	if user == (models.User{}) {
		return utils.ErrUserNotFound
	}
	token, err := utils.GenerateForgetPasswordToken(user)
	if err != nil {
		return err
	}

	resetLink = fmt.Sprintf("%s?token=%s", resetLink, token)
	err = utils.SendForgetPasswordEmail(user.FirstName, user.Email, resetLink)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ResetPassword(token string, password string) error {
	payload, err := utils.ValidateToken(token)
	if err != nil {
		return err
	}

	id := payload["id"].(float64)
	user, err := u.UserRepository.FindByID(int(id))
	if err != nil {
		return err
	}
	if user == (models.User{}) {
		return utils.ErrUserNotFound
	}
	password, _ = utils.HashPassword(password)
	user.Password = password
	_, err = u.UserRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ChangePassword(user models.User, oldPassword string, newPassword string) error {
	userObj, err := u.UserRepository.FindByID(user.Id)
	if err != nil {
		return err
	}
	if userObj == (models.User{}) {
		return utils.ErrUserNotFound
	}
	if !utils.ComparePassword(userObj.Password, oldPassword) {
		return utils.ErrWrongPassword
	}
	password, _ := utils.HashPassword(newPassword)
	user.Password = password
	_, err = u.UserRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) VerifyEmail(token string) error {
	payload, err := utils.ValidateToken(token)
	if err != nil {
		return err
	}

	id := payload["id"].(float64)
	user, err := u.UserRepository.FindByID(int(id))
	if err != nil {
		return err
	}
	if user == (models.User{}) {
		return utils.ErrUserNotFound
	}
	if user.Status {
		return utils.ErrUserAlreadyActive
	}
	user.Status = true
	_, err = u.UserRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}
