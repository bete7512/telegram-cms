package ports

import "github.com/bete7512/telegram-cms/models"

type UserService interface {
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id int) (bool, error)
	SignUp(user models.User) (models.User, error)
	Login(email string, password string) (models.User, error)
	VerifyEmail(token string) error
	ForgetPassword(email string) error
	ResetPassword(token string, password string) error
	ChangePassword(user models.User, oldPassword string, newPassword string) error
	ResendSignupVerificationEmail(email string) error
	ResendForgetPasswordEmail(email string) error

}

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id int) (bool, error)
	FindByEmail(email string) (models.User, error)
}