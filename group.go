package main

import (
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type Group struct {
    ID          string      `json:"id"`
    Name        string      `json:"name"`
    Description string      `json:"description"`
    IsPublic    bool        `json:"isPublic"`
    CreatedBy   string      `json:"createdBy"`
    UserList    []string    `json:"userList"`
    EventList   []string    `json:"eventList"`
}

var groups = make([]Group, 0)


func getAllPublicGroups(c *gin.Context) {
    publicGroups := make([]Group, 0)
    i := 0
    for i < len(groups) {
        if groups[i].IsPublic == true {
            publicGroups = append(publicGroups, groups[i])
        }
        i++
    }
    c.JSON(http.StatusOK, publicGroups)
}

func getGroupByID(c *gin.Context) {
    groupID := c.Param("groupID")
    group, err := getGroupByGroupID(groupID)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.IndentedJSON(http.StatusOK, group)
}

func getGroupsByUserID(c *gin.Context) {
    filteredGroups := make([]Group, 0)
    userID := c.Param("userID")
    for _, group := range groups {
        if group.CreatedBy == userID || userExistInUserList(group.UserList, userID) {
            filteredGroups = append(filteredGroups, group)
        }
    }
    c.JSON(http.StatusOK, filteredGroups)
}

func createGroup(c *gin.Context) {
    var newGroup Group
    
    if err := c.ShouldBindJSON(&newGroup); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newGroup.ID = uuid.New().String()
    groups = append(groups, newGroup)
    c.JSON(http.StatusCreated, newGroup)
}

func updateGroup(c *gin.Context) {
    var updatedGroup Group
    if err := c.ShouldBindJSON(&updatedGroup); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, group := range groups {
        if group.ID == updatedGroup.ID {
            groups[i] = updatedGroup
            c.JSON(http.StatusOK, groups[i])
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
}

func getGroupByGroupID(groupID string) (Group, error) {
    for _, group := range groups {
        if group.ID == groupID {
            return group, nil
        }
    }
    return Group{}, errors.New("Group not found")
}

func userExistInUserList(userList []string, userID string) bool {
    for _, id := range userList {
        if id == userID {
            return true
        }
    }
    return false
}
