package main

import (
	"net/http"

	"example.com/go-rest-api-sql-jwt/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//routes
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
}

func createEvent(c *gin.Context) {
	event := models.Event{}
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	event.ID = 1
	event.UserId = 1

	c.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}
