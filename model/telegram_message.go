package model

type TelegramMessage struct {
    Message struct {
        Text string `json:"text"`
    } `json:"message"`
    InlineQuery struct {
        Id    string `json:"id"`
        Query string `json:"query"`
    } `json:"inline_query"`
}
