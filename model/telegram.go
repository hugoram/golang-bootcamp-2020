package model

type TelegramUpdate struct {
    Message struct {
        Text string `json:"text"`
    } `json:"message"`
    InlineQuery struct {
        Id    string `json:"id"`
        Query string `json:"query"`
    } `json:"inline_query"`
}

type TelegramInlineKeyboardButton struct {
    Text string `json:"text"`
    URL  string `json:"url"`
}

type TelegramReplyMarkup struct {
    InlineKeyboard [][]TelegramInlineKeyboardButton `json:"inline_keyboard"`
}

type TelegramInlineResult struct {
    Type        string              `json:"type"`
    Id          string              `json:"id"`
    Title       string              `json:"title"`
    AudioURL    string              `json:"audio_url"`
    Performer   string              `json:"performer"`
    ReplyMarkup TelegramReplyMarkup `json:"reply_markup"`
}

type TelegramInlineResults struct {
    InlineQueryId string                 `json:"inline_query_id"`
    Results       []TelegramInlineResult `json:"results"`
}
