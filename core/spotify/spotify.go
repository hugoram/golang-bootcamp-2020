package spotify

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"

    "golang-bootcamp-2020/config"
    "golang-bootcamp-2020/model"
)

type SpotifyAuth struct {
    AccessToken string `json:"access_token"`
}

type SpotifySearchResult struct {
    Tracks struct {
        Items []model.SpotifyTrack `json:"items"`
    } `json:"tracks"`
}

func getToken() string {
    var auth SpotifyAuth

    client := &http.Client{}
    data := url.Values{}
    data.Set("grant_type", "client_credentials")
    req, _ := http.NewRequest("POST", config.SpotifyAuthURL, strings.NewReader(data.Encode()))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", "Basic "+config.SpotifyAuth)
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    jsonErr := json.NewDecoder(strings.NewReader(string(body))).Decode(&auth)
    if jsonErr != nil {
        log.Println(jsonErr)
    }

    return auth.AccessToken
}

func GetTracks(search string) []model.SpotifyTrack {
    var searchResult SpotifySearchResult
    token := getToken()

    client := &http.Client{}
    req, _ := http.NewRequest("GET", config.SpotifyApiURL+"/search?type=track&q="+url.QueryEscape(search), nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    jsonErr := json.NewDecoder(strings.NewReader(string(body))).Decode(&searchResult)
    if jsonErr != nil {
        log.Println(jsonErr)
    }

    return searchResult.Tracks.Items
}
