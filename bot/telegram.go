package bot

import (
	"fmt"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/httpx"
)

// API Document: https://core.telegram.org/bots/api#sendmessage

type Telegram struct {
	token  string
	chatId string
}

var _ Interface = (*Telegram)(nil)

func NewTelegram(token, chatId string) *Telegram {
	return &Telegram{token: token, chatId: chatId}
}

func (bot *Telegram) Name() string {
	return "telegram"
}

const TelegramAddr = "https://api.telegram.org"

type TelegramReq struct {
	ChatId                   string `json:"chat_id"`
	MessageThreadId          int    `json:"message_thread_id,omitempty"`
	Text                     string `json:"text"`
	ParseMode                string `json:"parse_mode,omitempty"`
	Entities                 any    `json:"entities,omitempty"`
	DisableWebPagePreview    bool   `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool   `json:"disable_notification,omitempty"`
	ProtectContent           bool   `json:"protect_content,omitempty"`
	ReplyToMessageId         int    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              any    `json:"reply_markup,omitempty"`
}

type TelegramResp struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func (bot *Telegram) Send(text string) error {
	return bot.SendRaw(&TelegramReq{ChatId: bot.chatId, Text: text})
}

func (bot *Telegram) SendRaw(raw any) error {
	addr := TelegramAddr + "/bot" + bot.token + "/sendMessage"
	resp, err := httpx.R().
		SetHeader("Content-Type", "application/json").
		SetBody(json.MustMarshalString(raw)).
		SetResult(&TelegramResp{}).
		Post(addr)
	if err != nil {
		return fmt.Errorf("bot: telegram: %w", err)
	}
	if e, ok := resp.Result().(*TelegramResp); ok && !e.Ok {
		return fmt.Errorf("bot: telegram: %d, %s", e.ErrorCode, e.Description)
	}
	return nil
}
