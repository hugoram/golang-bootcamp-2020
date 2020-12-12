package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Could not load .env file")
    }

    telegramUrl := "https://api.telegram.org/bot" + os.Getenv("TOKEN") + "/setWebhook"
    botUrl := os.Getenv("APP_URL") + "/telegram"

    reqBody, err := json.Marshal(map[string]string{
        "url": botUrl,
    })
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post(telegramUrl, "application/json", bytes.NewBuffer(reqBody))
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
