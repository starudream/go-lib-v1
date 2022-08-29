package main

import (
	"github.com/starudream/go-lib/bot"
)

func main() {
	e1 := bot.Send("hello")
	if e1 != nil {
		panic(e1)
	}

	e2 := bot.SendRaw(&bot.DingtalkReq{MsgType: "markdown", Markdown: &bot.DingtalkMarkdown{Title: "foo", Text: "bar"}})
	if e2 != nil {
		panic(e2)
	}
}
