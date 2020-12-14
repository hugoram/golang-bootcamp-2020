package controller

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"

    "golang-bootcamp-2020/core/spotify"
    "golang-bootcamp-2020/core/telegram"
    "golang-bootcamp-2020/model"
)

func Telegram(context *gin.Context) {
    var telegramUpdate model.TelegramUpdate

    if err := context.ShouldBindJSON(&telegramUpdate); err != nil {
        log.Println(err)
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if telegramUpdate.InlineQuery.Query == "" {
        context.JSON(200, nil)
        return
    }

    tracks := spotify.GetTracks(telegramUpdate.InlineQuery.Query)

    var results []model.TelegramInlineResult

    for _, track := range tracks {
        if len(track.Artists) == 0 || track.PreviewURL == "" {
            continue
        }

        result := model.TelegramInlineResult{
            "audio",
            track.Id,
            track.Name,
            track.PreviewURL,
            track.Artists[0].Name,
            model.TelegramReplyMarkup{
                [][]model.TelegramInlineKeyboardButton{{{
                    "Open in Spotify",
                    track.ExternalURLs.Spotify,
                }}},
            },
        }

        results = append(results, result)
    }

    defer telegram.SendInlineResults(model.TelegramInlineResults{
        telegramUpdate.InlineQuery.Id,
        results,
    })

    context.JSON(200, nil)
}
