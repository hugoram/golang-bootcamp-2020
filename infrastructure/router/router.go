package router

import (
    "github.com/gin-gonic/gin"

    "golang-bootcamp-2020/config"
    "golang-bootcamp-2020/controller"
)

func NewRouter(engine *gin.Engine) {
    engine.POST("/"+config.TelegramToken, controller.Telegram)
}
