package main

import (
    "time"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type User struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    Gender      string `json:"gender"`
    DateOfBirth time.Time `json:"dateOfBirth"`
}

var users = make([]User, 0)

func getUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
    userID := c.Param("userID")
    for _, user := range users {
        if user.ID == userID {
            c.JSON(http.StatusOK, user)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func createUser(c *gin.Context) {
    var newUser User
    
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if newUser.Gender != "M" && newUser.Gender != "F" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Gender type"})
        return
    }

    newUser.ID = uuid.New().String()
    users = append(users, newUser)
    c.JSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
    var updatedUser User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, user := range users {
        if user.ID == updatedUser.ID {
            users[i].Name = updatedUser.Name
            users[i].Email = updatedUser.Email
            users[i].Gender = updatedUser.Gender
            users[i].DateOfBirth = updatedUser.DateOfBirth
            c.JSON(http.StatusOK, users[i])
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func validateUserMiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.Param("userID")
        if userID == "" || validateUser(userID) {
            c.Next()
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        }
    }
}

func validateUser(userID string) bool {
    isValidUser := false
    for _, user := range users {
        if user.ID == userID {
            isValidUser = true
            break
        }
    }
    return isValidUser
}