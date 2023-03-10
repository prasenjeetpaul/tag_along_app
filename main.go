package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)


func main() {
    fmt.Println("Welcome to Tag Along App")
    router := gin.Default()

    userRouter := router.Group("/users")
    {
        userRouter.GET("", getUsers)
        userRouter.GET("/:userID", getUserByID)
        userRouter.POST("", createUser)
        userRouter.PUT("", updateUser)
    }

    groupRouter := router.Group("/groups")
    {
        groupRouter.GET("", getAllPublicGroups)
        groupRouter.GET("/:groupID",getGroupByID)
        groupRouter.GET("/user/:userID", getGroupsByUserID)
        groupRouter.POST("", createGroup)
        groupRouter.PUT("", updateGroup)
    }

    router.Run()
}
