package config

import "os"

var AppURL = os.Getenv("APP_URL")
var SpotifyApiURL = "https://api.spotify.com/v1"
var SpotifyAuth = os.Getenv("SPOTIFY_AUTH")
var SpotifyAuthURL = "https://accounts.spotify.com/api/token"
var TelegramToken = os.Getenv("TELEGRAM_TOKEN")
var TelegramApiURL = "https://api.telegram.org/bot" + TelegramToken
