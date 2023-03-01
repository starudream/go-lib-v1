package bot

import (
	"github.com/starudream/go-lib/config"

	"github.com/starudream/go-lib/internal/ilog"
)

type Interface interface {
	Name() string
	Send(text string) error
	SendRaw(raw any) error
}

var _b Interface

func init() {
	// dingtalk
	{
		token, secret := config.GetString("dingtalk.token"), config.GetString("dingtalk.secret")
		if token != "" && secret != "" {
			_b = NewDingtalk(token, secret)
			ilog.X.Debug().Msgf("registry default bot as dingtalk")
		}
	}

	// telegram
	{
		token, chatId := config.GetString("telegram.token"), config.GetString("telegram.chat_id")
		if token != "" && chatId != "" {
			_b = NewTelegram(token, chatId)
			ilog.X.Debug().Msgf("registry default bot as telegram")
		}
	}
}

func Init(b Interface) {
	_b = b
}

func Bot() Interface {
	return _b
}

func Send(text string) error {
	if _b == nil {
		return nil
	}
	return _b.Send(text)
}

func SendRaw(raw any) error {
	if _b == nil {
		return nil
	}
	return _b.SendRaw(raw)
}
