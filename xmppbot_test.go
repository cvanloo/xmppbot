package xmppbot_test

import (
	"xmppbot"
	"xmppbot/cmd"
)

func ExampleUsage() {
	room := xmppbot.Target{}.Room("bots@conference.example.com")
	bot := xmppbot.New().Login("username", "password").Join(room)

	bot.ListenToCommand(xmppbot.Global("echo", func(req xmppbot.Message) {
		resp := xmppbot.Message{}.To(req.Target).Text(req.Text).Tag(req.From)
		bot.SendTextMessage(resp)
	}))
	bot.ListenToCommand(xmppbot.Tagged("ping", func(req xmppbot.Message) {
		resp := xmppbot.Message{}.To(req.Target).Text("pong").Tag(req.From).SendFrom(bot)
	}))

	hello := xmppbot.Message{}.To(room).Text("bot ready")
	bot.SendMessage(hello)
}
