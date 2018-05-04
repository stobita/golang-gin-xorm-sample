package handler

import "github.com/gin-gonic/gin"

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Hello!")
	}
}
