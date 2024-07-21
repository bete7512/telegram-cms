package handlers

import (
	"log"
	"strconv"

	"github.com/bete7512/telegram-cms/models"
	"github.com/bete7512/telegram-cms/services"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	services.UserService
}

func NewUserHandlers(userService services.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

//	@Summary		Get all users
//	@Schemes		http
//	@Description	Get all users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.User
//	@Security		Bearer
//	@Router			/users	[get]
func (h *UserHandlers) GetAllUsers(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	log.Println(user)
	if user.Status != false {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	users, err := h.UserService.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

//	@Summary		Get user by id
//	@Schemes		http
//	@Description	Get user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	message	"Get user by id"
//	@Security		Bearer
//	@Router			/users/{id}	[get]
func (h *UserHandlers) GetUserByID(ctx *gin.Context) {
	// id := ctx.Param("id")
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.UserService.FindByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}

//	@Summary		Create user
//	@Schemes		http
//	@Description	Create user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		201	{string}	message	"User created"
//	@Security		ApiKeyAuth
//	@Router			/users	[post]
func (h *UserHandlers) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newUser, err := h.UserService.Create(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, newUser)
}

//	@Summary		Update user
//	@Schemes		http
//	@Description	Update user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	message	"User updated"
//	@Security		ApiKeyAuth
//	@Router			/users	[put]
func (h *UserHandlers) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := h.UserService.Update(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, updatedUser)
}

//	@Summary		Change password
//	@Schemes		http
//	@Description	Change password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.ChangePasswordRequest	true	"User"
//	@Success		200		{object}	models.ChangePasswordResponse
//	@Security		Bearer
//	@Router			/user/change-password/	[post]
func (h *UserHandlers) ChangePassword(ctx *gin.Context) {
	var changePasswordRequest models.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&changePasswordRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := ctx.MustGet("user").(models.User)
	err := h.UserService.ChangePassword(user,changePasswordRequest.OldPassword, changePasswordRequest.NewPassword)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Password updated"})
}
//	@Summary		Delete user
//	@Schemes		http
//	@Description	Delete user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200			{string}	message	"User deleted"
//	@Router			/users/{id}	[delete]
func (h *UserHandlers) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	deleted, err := h.UserService.Delete(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if deleted {
		ctx.JSON(200, gin.H{"message": "User deleted successfully"})
	} else {
		ctx.JSON(500, gin.H{"error": "User could not be deleted"})
	}
}
