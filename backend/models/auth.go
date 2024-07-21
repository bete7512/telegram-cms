package models

type SignupRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RedirectUri string `json:"redirect_uri" default:"http://localhost:8767"`
}

type SignupResponse struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`

}

type VerifyEmailRequest struct {
	Token string `json:"token"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponse struct {
	Message string `json:"message"`
}

type VerifyEmailResponse struct {
	Message string `json:"message"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}

type ForgetPasswordResponse struct {
	Message string `json:"message"`
}
