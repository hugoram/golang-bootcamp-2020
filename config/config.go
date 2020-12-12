package config

import "os"

var AppURL = os.Getenv("APP_URL")
var TelegramToken = os.Getenv("TELEGRAM_TOKEN")
var TelegramApiURL = "https://api.telegram.org/bot" + TelegramToken
