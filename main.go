package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/ping", func(ctx *gin.Context) {
        ctx.String(http.StatusOK, "HELLO WORLD")
    })

    return r
}

func main() {
    r := setupRouter()
    r.Run(":8080")
}