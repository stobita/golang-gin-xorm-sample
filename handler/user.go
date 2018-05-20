package handler

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-gin-xorm-sample/models"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "secret"

func Users() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := models.User{}.GetAll()
		c.JSON(200, result)
	}
}

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		result := models.User{}.GetByID(id)
		c.JSON(200, result)
	}
}

type SignInForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json SignInForm
		c.ShouldBindJSON(&json)
		name := json.Name
		password := json.Password
		if name == "" || password == "" {
			c.JSON(400, "value not")
			return
		}
		result := models.User{Name: name}.GetByName()
		if result == nil {
			c.JSON(400, "not find user")
		}
		err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
		if err == nil {
			token := jwt.New(jwt.GetSigningMethod("HS256"))
			token.Claims = jwt.MapClaims{
				"user": name,
				"exp":  time.Now().Add(time.Hour * 1).Unix(),
			}
			tokenString, jwtError := token.SignedString([]byte(secretKey))
			if jwtError == nil {
				c.JSON(200, gin.H{"token": tokenString})
			} else {
				c.JSON(400, "could not generate token")
			}
		} else {
			c.JSON(400, err)
		}
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		confirmationPassword := c.PostForm("confirmationPassword")
		if name == "" || password == "" || confirmationPassword == "" {
			c.JSON(400, "value not")
			return
		}
		if password != confirmationPassword {
			c.JSON(400, "password missmatch")
		}
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, err)
			return
		}
		result := models.User{Name: name, Password: string(encryptedPassword)}.Insert()
		c.JSON(200, result)
	}
}
