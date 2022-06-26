package event

import "github.com/gin-gonic/gin"

var Controller EventController = &event{}

type EventController interface {
	GetAllEvents(ctx *gin.Context)
	GetEventByID(ctx *gin.Context)
}

type event struct{}
