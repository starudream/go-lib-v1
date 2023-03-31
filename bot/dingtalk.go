package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/httpx"
)

// API Document: https://open.dingtalk.com/document/group/custom-robot-access

type Dingtalk struct {
	token  string
	secret string
}

var _ Interface = (*Dingtalk)(nil)

func NewDingtalk(token, secret string) *Dingtalk {
	return &Dingtalk{token: token, secret: secret}
}

func (bot *Dingtalk) Name() string {
	return "dingtalk"
}

const DingtalkAddr = "https://oapi.dingtalk.com/robot/send"

type DingtalkReq struct {
	MsgType string      `json:"msgtype"` // such as: "text", "markdown"
	At      *DingtalkAt `json:"at,omitempty"`

	Text     *DingtalkText     `json:"text,omitempty"`
	Markdown *DingtalkMarkdown `json:"markdown,omitempty"`
}

type DingtalkAt struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type DingtalkText struct {
	Content string `json:"content"`
}

type DingtalkMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DingtalkResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (bot *Dingtalk) Send(text string) error {
	return bot.SendRaw(&DingtalkReq{MsgType: "text", Text: &DingtalkText{Content: text}})
}

func (bot *Dingtalk) SendRaw(raw any) error {
	addr := DingtalkAddr + "?access_token=" + bot.token
	if bot.secret != "" {
		ts, sign := bot.Sign(bot.secret)
		addr += "&timestamp=" + ts + "&sign=" + sign
	}
	resp, err := httpx.R().
		SetHeader("Content-Type", "application/json").
		SetBody(json.MustMarshalString(raw)).
		SetResult(&DingtalkResp{}).
		Post(addr)
	if err != nil {
		return fmt.Errorf("bot: dingtalk: %w", err)
	}
	if e, ok := resp.Result().(*DingtalkResp); ok && e.ErrCode != 0 {
		return fmt.Errorf("bot: dingtalk: %d, %s", e.ErrCode, e.ErrMsg)
	}
	return nil
}

func (bot *Dingtalk) Sign(secret string) (string, string) {
	milli := strconv.FormatInt(time.Now().UnixMilli(), 10)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(milli + "\n" + secret))
	return milli, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
