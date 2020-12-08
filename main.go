package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "golang-bootcamp-2020/infrastructure/router"
)

func loadEnvVars() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Could not load .env file")
    }
}

func main() {
    loadEnvVars()
    server := gin.Default()
    router.NewRouter(server)
    server.Run()
}
