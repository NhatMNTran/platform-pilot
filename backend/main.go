package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Environment struct {
    ID     string `json:"id"`
    Status string `json:"status"`
}

func main() {
    router := gin.Default()

    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    router.GET("/environments/:id", func(c *gin.Context) {
        env := Environment{
            ID:     c.Param("id"),
            Status: "READY",
        }

        c.JSON(http.StatusOK, env)
    })

    router.GET("/metrics", gin.WrapH(promhttp.Handler()))

    router.Run(":8080")
}