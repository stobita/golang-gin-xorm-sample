package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/gin-xorm-sample/app/models"
)

func Users() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := models.User{}.GetAll()
		c.JSON(200, result)
	}
}
