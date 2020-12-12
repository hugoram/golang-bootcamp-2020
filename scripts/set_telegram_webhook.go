package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"

    _ "github.com/joho/godotenv/autoload"

    "golang-bootcamp-2020/config"
)

func main() {
    reqBody, err := json.Marshal(map[string]string{
        "url": config.AppURL + "/" + config.TelegramToken,
    })
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post(config.TelegramApiURL+"/setWebhook", "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(string(body))
}
