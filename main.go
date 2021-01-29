package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/joho/godotenv/autoload"

    "golang-bootcamp-2020/infrastructure/router"
)

func main() {
    server := gin.Default()
    router.NewRouter(server)
    server.Run()
}
