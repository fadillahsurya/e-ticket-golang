package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()

	r := gin.Default()

	r.POST("/login", Login)
	r.POST("/terminals", AuthMiddleware(), CreateTerminal)

	r.Run(":8080")
}
