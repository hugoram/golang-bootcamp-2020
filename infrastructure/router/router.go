package router

import (
    "github.com/gin-gonic/gin"
)

func telegramHandler(context *gin.Context) {
    var data = []int{1, 2, 3}
    context.JSON(200, data)
}

func NewRouter(engine *gin.Engine) {
    engine.GET("/telegram", telegramHandler)
}
