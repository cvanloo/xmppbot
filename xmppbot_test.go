package xmppbot_test

import (
	"log"
	"github.com/cvanloo/xmppbot"
)

func ExampleUsage() {
	room := xmppbot.Target{}.Room("bots@conference.example.com")
	bot := xmppbot.New().Login("username@example.com", "password").Join(room.ToTarget())
	if bot.Error != nil {
		log.Fatal(bot.Error)
	}

	bot.ListenToCommand(xmppbot.Global("echo", func(req xmppbot.Message) {
		resp := xmppbot.Message{}.To(req.Target).Text(req.Content).Tag(req.From)
		bot.SendTextMessage(resp.Content)
	}))
	bot.ListenToCommand(xmppbot.Tagged("ping", func(req xmppbot.Message) {
		xmppbot.Message{}.To(req.Target).Text("pong").Tag(req.From).SendFrom(bot.ToTarget())
	}))

	hello := xmppbot.Message{}.To(room.ToTarget()).Text("bot ready")
	bot.SendMessage(hello)
}
