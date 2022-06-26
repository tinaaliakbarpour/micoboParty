package event

import (
	"fmt"
	eventrepo "micobianParty/domain/repository/event"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetAllEvents will return a list of all events in micobo
func (event) GetAllEvents(ctx *gin.Context) {
	events, err := eventrepo.Repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed fetching events with error : " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "fetched events successfully",
		"events":  events,
	})
}

//GetEventByID will return an event by specific id
func (event) GetEventByID(ctx *gin.Context) {
	input := ctx.Param("event_id")

	id, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed converting string to uint with error : " + err.Error(),
		})
		return
	}

	fetchedEvent, err := eventrepo.Repository.Get(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed fetching event with id %d with error : %s ", id, err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "fetched  event successfully",
		"event":   fetchedEvent,
	})
}
