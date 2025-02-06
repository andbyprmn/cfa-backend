package handler

import (
	"cfa-backend/auth"
	"cfa-backend/helper"
	"cfa-backend/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

// RegisterUser godoc
// @Summary      Register a new user
// @Description  This endpoint registers a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body  body  user.RegisterUserInput  true  "User data"
// @Success      200   {object}  user.UserFormatter
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /users [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Registered account failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Registered account failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Registered account failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// LoginUser godoc
// @Summary      Login an user
// @Description  This endpoint login an users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body  body  user.LoginInput  true  "User login data"
// @Success      200   {object}  user.UserFormatter
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /sessions [post]
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, errToken := h.authService.GenerateToken(loggedinUser.ID)
	if errToken != nil {
		errorMessage := gin.H{"errors": errToken.Error()}

		response := helper.APIResponse("Login failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Successfuly logged in!", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

// CheckEmailAvailabilityRegister godoc
// @Summary      Check email availability register a new user
// @Description  This endpoint check email availability register a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body  body  user.CheckEmailInput  true  "Email availability register"
// @Success      200   {object}  user.UserFormatter
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /email_checkers [post]
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Email checking failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email is available!"
	} else {
		metaMessage = "Email has been registered!"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

// UploadAvatar godoc
// @Summary      Upload user avatar
// @Description  Upload a new avatar for the user
// @Tags         Users
// @Accept       multipart/form-data
// @Produce      json
// @Param        avatar  formData  file  true  "Avatar image file"
// @Success      200   {object}  helper.Response
// @Failure      400   {object}  helper.Response
// @Failure      422   {object}  helper.Response
// @Router       /avatars [post]
func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded!", http.StatusOK, "error", data)

	c.JSON(http.StatusOK, response)
}
