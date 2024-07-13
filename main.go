package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg": "Hello World",
        })
    })

    r.Run(":8080") // Exposes the server on port 8080
}
