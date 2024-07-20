package handlers

import (
	"github.com/bete7512/telegram-cms/models"
	"github.com/bete7512/telegram-cms/services"
	"github.com/bete7512/telegram-cms/utils"
	"github.com/gin-gonic/gin"
)

type AuthenticationHandlers struct {
	AuthenticationsService services.UserService
}

func NewAuthenticationHandlers(authenticationsService services.UserService) *AuthenticationHandlers {
	return &AuthenticationHandlers{AuthenticationsService: authenticationsService}
}

// @Summary		Sign up
// @Schemes		http
// @Description	Sign up
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Param			user	body		models.SignupRequest	true	"User"
// @Success		200		{object}	models.SignupResponse
// @Router			/signup	[post]
func (h *AuthenticationHandlers) SignUp(ctx *gin.Context) {

	var signupRequest models.SignupRequest
	if err := ctx.ShouldBindJSON(&signupRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newUser, err := h.AuthenticationsService.SignUp(signupRequest)
	if err != nil {
		code, err := utils.FilterError(err)
		ctx.JSON(code, gin.H{"error": err})
		return
	}
	signupResponse := models.SignupResponse{
		Id:        newUser.Id,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}
	err = utils.SendSignupEmail(signupRequest.FirstName, newUser.Email)
	if err != nil {
		code, err := utils.FilterError(err)
		ctx.JSON(code, gin.H{"error": err})
		return
	}
	ctx.JSON(201, signupResponse)
}

// @Summary		Login
// @Schemes		http
// @Description	Login
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Param			user	body		models.LoginRequest	true	"User"
// @Success		200							{object}	models.LoginResponse
// @Router			/login/	[post]
func (h *AuthenticationHandlers) Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.AuthenticationsService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"token": token})
}

// @Summary		Forget password
// @Schemes		http
// @Description	Forget password
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Success		200							{string}	message	"Email sent"
// @Router			/forget-password/{email}	[post]
func (h *AuthenticationHandlers) ForgetPassword(ctx *gin.Context) {
	email := ctx.Param("email")
	err := h.AuthenticationsService.ForgetPassword(email)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Email sent"})
}

// @Summary		Reset password
// @Schemes		http
// @Description	Reset password
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Success		200									{string}	message	"Password updated"
// @Router			/reset-password/{token}/{password}	[post]
func (h *AuthenticationHandlers) ResetPassword(ctx *gin.Context) {
	// TODO: implement reset password logic
	token := ctx.Param("token")
	password := ctx.Param("password")
	err := h.AuthenticationsService.ResetPassword(token, password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Password updated"})
}

// @Summary		Change password
// @Schemes		http
// @Description	Change password
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Success		200												{string}	message	"Password updated"
// @Router			/change-password/{oldPassword}/{newPassword}	[post]
func (h *AuthenticationHandlers) ChangePassword(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	oldPassword := ctx.Param("oldPassword")
	newPassword := ctx.Param("newPassword")
	err := h.AuthenticationsService.ChangePassword(user, oldPassword, newPassword)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Password updated"})
}

// @Summary		Verify email
// @Schemes		http
// @Description	Verify email
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Success		200						{string}	message	"Email verified"
// @Router			/verify-email/{token}	[post]
func (h *AuthenticationHandlers) VerifyEmail(ctx *gin.Context) {
	token := ctx.Param("token")
	err := h.AuthenticationsService.VerifyEmail(token)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Email verified"})
}
