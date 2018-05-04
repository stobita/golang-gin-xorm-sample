package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-gin-xorm-sample/handler"
)

func main() {
	router := gin.Default()
	router.GET("/", handler.Hello())

	router.Run(":8080")
}
