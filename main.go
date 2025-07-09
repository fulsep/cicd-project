package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Backend is running well",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "List all users",
			"results": []map[string]string{
				{
					"name": "John",
				},
				{
					"name": "Ed",
				},
				{
					"name": "Jane",
				},
			},
		})
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}
