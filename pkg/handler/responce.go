package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(g *gin.Context, stausCode int, message string) {
	logrus.Errorf(message)
	g.AbortWithStatusJSON(stausCode, errorResponse{message})
}
