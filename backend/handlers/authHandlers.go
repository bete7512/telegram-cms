package handlers

import (
	"net/http"

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

//	@Summary		Sign up
//	@Schemes		http
//	@Description	Sign up
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user			body		models.SignupRequest	true	"User"
//	@Success		200				{object}	models.SignupResponse
//	@Router			/auth/signup	[post]
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
	// TODO: generate signup token
	signup_token, err := utils.GenerateSignupToken(newUser)
	if err != nil {
		code, err := utils.FilterError(err)
		ctx.JSON(code, gin.H{"error": err})
		return
	}

	redirectUri := signupRequest.RedirectUri + "/api/v1/auth/verify-email" + "?token=" + signup_token
	err = utils.SendSignupEmail(signupRequest.FirstName, newUser.Email, redirectUri)
	if err != nil {
		code, err := utils.FilterError(err)
		ctx.JSON(code, gin.H{"error": err})
		return
	}
	ctx.JSON(201, signupResponse)
}

//	@Summary		Login
//	@Schemes		http
//	@Description	Login
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user			body		models.LoginRequest	true	"User"
//	@Success		200				{object}	models.LoginResponse
//	@Router			/auth/login/	[post]
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

//	@Summary		Forget password
//	@Schemes		http
//	@Description	Forget password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user					body		models.ForgetPasswordRequest	true	"User"
//	@Success		200						{object}	models.ForgetPasswordResponse
//	@Router			/auth/forget-password/	[post]
func (h *AuthenticationHandlers) ForgetPassword(ctx *gin.Context) {
	body := models.ForgetPasswordRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resetLink := "http://localhost:8767/auth/reset-password"
	err := h.AuthenticationsService.ForgetPassword(body.Email, resetLink)
	if err != nil {
		code, err := utils.FilterError(err)
		ctx.JSON(code, gin.H{"error": err})
		return
	}
	ctx.JSON(200, models.ForgetPasswordResponse{Message: "Email sent"})
}

//	@Summary		Reset password
//	@Schemes		http
//	@Description	Reset password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user					body		models.ResetPasswordRequest	true	"User"
//	@Success		200						{object}	models.ResetPasswordResponse
//	@Router			/auth/reset-password/	[post]
func (h *AuthenticationHandlers) ResetPassword(ctx *gin.Context) {
	resetPasswordRequest := models.ResetPasswordRequest{}
	if err := ctx.ShouldBindJSON(&resetPasswordRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.AuthenticationsService.ResetPassword(resetPasswordRequest.Token, resetPasswordRequest.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Password updated successfully"})
}

//	@Summary		Verify email
//	@Schemes		http
//	@Description	Verify email
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200					{string}	message	"Email verified"
//
//	@Param			token				query		string	true	"Verification token"
//
//	@Router			/auth/verify-email	[get]
func (h *AuthenticationHandlers) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query("token")
	err := h.AuthenticationsService.VerifyEmail(token)
	frontendURL := "https://localhost:3000/verified"
	if err != nil {
		if err.Error() == utils.ErrUserAlreadyActive.Error() {
			ctx.Redirect(http.StatusFound, frontendURL)
		}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, frontendURL)
}
