package telegram

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"

    "golang-bootcamp-2020/config"
    "golang-bootcamp-2020/model"
)

func SendInlineResults(results model.TelegramInlineResults) {
    reqBody, err := json.Marshal(results)
    if err != nil {
        log.Println(err)
    }

    resp, err := http.Post(config.TelegramApiURL+"/answerInlineQuery", "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        log.Println(err)
    }

    defer resp.Body.Close()
}
