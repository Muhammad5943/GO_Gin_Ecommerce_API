package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/muhammad5943/go_gin_ecommerce_api/dtos"
	"github.com/muhammad5943/go_gin_ecommerce_api/models"
	"github.com/muhammad5943/go_gin_ecommerce_api/services"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

func UsersRegistration(c *gin.Context) {

	var json dtos.RegisterRequestDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err := services.CreateOne(&models.User{
		Username:  json.Username,
		Password:  string(password),
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Email:     json.Email,
	}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success":       true,
		"full_messages": []string{"User created successfully"}})
}

func UsersLogin(c *gin.Context) {

	var json dtos.LoginRequestDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}

	user, err := services.FindOneUser(&models.User{Username: json.Username})

	if err != nil {
		c.JSON(http.StatusForbidden, dtos.CreateDetailedErrorDto("login_error", err))
		return
	}

	if user.IsValidPassword(json.Password) != nil {
		c.JSON(http.StatusForbidden, dtos.CreateDetailedErrorDto("login", errors.New("invalid credentials")))
		return
	}

	c.JSON(http.StatusOK, dtos.CreateLoginSuccessful(&user))

}
