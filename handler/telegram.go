package handler

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"

    "golang-bootcamp-2020/model"
)

func TelegramHandler(context *gin.Context) {
    var json model.TelegramMessage

    if err := context.ShouldBindJSON(&json); err != nil {
        log.Println(err)
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Println(json)

    context.JSON(200, nil)
}
