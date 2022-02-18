package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorN struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}


func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorN{message})
}