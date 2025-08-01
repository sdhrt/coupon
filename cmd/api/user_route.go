package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// handles /register route
//
//	 required variables for the body
//		 name string
//		 email string
//		 password string
func (app *application) register_user(c *gin.Context) {
	var user_credentials struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&user_credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request when parsing body",
		})
		return
	}

	_, err = app.models.Users.Create_user(user_credentials.Name, user_credentials.Email, user_credentials.Password)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "The email address is already associated with an account") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Email address already exists",
			})
			return

		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Registered user successfully",
	})
}

// handles /login route
//
//	 required arguments for the request
//		 email string
//		 password string
func (app *application) login_user(c *gin.Context) {
	var user_credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&user_credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request when parsing body",
		})
		return
	}

	user, err := app.models.Users.Validate_user(user_credentials.Email, user_credentials.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}
	access_token, err := app.models.Users.Get_access_token(user, app.config.jwt_secret)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.Set("user", user)
	c.JSON(http.StatusOK, gin.H{
		"message":      "access token has been provided",
		"access_token": access_token,
	})
}
