package main

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-gin-xorm-sample/handler"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/", handler.Hello())
	router.POST("/signin", handler.SignIn())

	authorized := router.Group("/")
	authorized.Use(TokenAuthenticationMiddleware())
	{
		authorized.GET("/users", handler.Users())
		authorized.GET("/user/:id", handler.User())
		authorized.POST("/signup", handler.SignUp())
	}

	// for live reload
	port := os.Getenv("PORT")
	router.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

type AuthenticationToken struct {
	Token string `json:"token"`
}

var secretKey = "secret"

func TokenAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := string(c.GetHeader("api-key"))

		if tokenString == "" {
			c.JSON(400, "tokenString is necessary")
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// check signin method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(secretKey), nil
		})
		if err != nil {
			c.JSON(400, "token error")
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["user"] == nil {
				c.JSON(400, "no user")
				c.Abort()
				return
			}
		} else {
			c.JSON(400, "claim error")
			c.Abort()
			return
		}
	}
}
