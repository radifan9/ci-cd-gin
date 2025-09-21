package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to backend")
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": []struct {
				Id   int
				Name string
			}{
				{Id: 1, Name: "Slamet"},
				{Id: 2, Name: "Darari"},
				{Id: 3, Name: "Ucuuup"},
			},
		})
	})

	r.Run(":8080")
	log.Println("Backend is Running")
}
