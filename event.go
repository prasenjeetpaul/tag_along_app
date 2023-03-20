package main

import (
    "time"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type Event struct {
    ID                      string      `json:"id"`
    Name                    string      `json:"name"`
    Description             string      `json:"description"`
    CreatedBy               string      `json:"createdBy"`
    StartTime               time.Time   `json:"startTime"`
    Duration                int         `json:"duration"`
    InvitedUsers            []string    `json:"invitedUsers"`
    AcceptedUsers           []string    `json:"acceptedUsers"`
    AcceptanceRatio         float64     `json:"acceptanceRatio"`
}

var events = make([]Event, 0)

func getEventByID( c *gin.Context) {
    eventID := c.Param("eventID")
    for _, event := range events {
        if event.ID == eventID {
            c.JSON(http.StatusOK, event)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
}

func getEventByGroupID( c *gin.Context) {
    groupID := c.Param("groupID")
    filteredEvents := make([]Event, 0)
    for _, event := range events {
        if event.CreatedBy == groupID {
            filteredEvents = append(filteredEvents, event)
        }
    }
    c.JSON(http.StatusOK, filteredEvents)
}

func getEventByUserID( c *gin.Context) {
    userID := c.Param("userID")
    filteredEvents := make([]Event, 0)
    for _, event := range events {
        if userExistInUserList(event.InvitedUsers, userID) || userExistInUserList(event.AcceptedUsers, userID) {
            filteredEvents = append(filteredEvents, event)
        }
    }
    c.JSON(http.StatusOK, filteredEvents)
}

// TODO: validate user ids in invited/accepted users list before insert
func createEvent(c *gin.Context) {
    var newEvent Event
    
    if err := c.ShouldBindJSON(&newEvent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newEvent.ID = uuid.New().String()
    events = append(events, newEvent)
    c.JSON(http.StatusCreated, newEvent)
}


func updateEvent(c *gin.Context) {
    var updatedEvent Event
    if err := c.ShouldBindJSON(&updatedEvent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, event := range events {
        if event.ID == updatedEvent.ID {
            events[i] = updatedEvent
            c.JSON(http.StatusOK, events[i])
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
}
