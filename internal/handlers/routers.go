package handlers

import "github.com/gin-gonic/gin"

func Manejadores() {
	router := gin.Default()

	router.GET("/ping", ping2)

	router.Run(":8080")
}
