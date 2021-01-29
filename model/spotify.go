package model

type SpotifyTrack struct {
    Id         string `json:"id"`
    Name       string `json:"name"`
    PreviewURL string `json:"preview_url"`
    Artists    []struct {
        Name string `json:"name"`
    } `json:"artists"`
    ExternalURLs struct {
        Spotify string `json:"spotify"`
    } `json:"external_urls"`
}
