package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-gin-xorm-sample/models"
)

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

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		result := models.User{}.Insert(name)
		c.JSON(200, result)
	}
}
