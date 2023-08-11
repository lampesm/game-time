package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lampesm/game-time/db"
	"github.com/lampesm/game-time/forms"
)

func CreateGame(c *gin.Context) {
	var payload forms.GamePayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.InsertGame(payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":  false,
		"status": http.StatusCreated,
		"msg":    "ok",
	})
}
