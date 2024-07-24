package xmppbot

import (
	"context"
	"log" // @todo: slog?

	"mellium.im/xmpp"
	"mellium.im/xmpp/jid"
)

func New() *Bot {
	xmpp.DialClientSession(
		context.Background(),
		jid.MustParse(""),
	)
	return nil
}

type Bot struct{
	Target
	Error error
}

func (b *Bot) reportError(err error) {
	if err != nil {
		b.Error = err
		log.Printf("xmppbot: %s", b.Error)
	}
}

func (b *Bot) Login(user, pass string) *Bot {
	return b
}

func (b *Bot) Join(t Target) *Bot {
	return b
}

func (b *Bot) SendMessage(msg Message) {
}

func (b *Bot) ListenToCommand(f ListenFunc) {
}

func (b *Bot) SendTextMessage(text string) {
}

type (
	Target struct{}
	RoomTarget struct{
		Target
	}
	ContactTarget struct{
		Target
	}
)

func (t Target) ToTarget() Target {
	return t
}

func (t Target) Room(room string) RoomTarget {
	return RoomTarget{t}
}

func (t Target) Contact(contact string) ContactTarget {
	return ContactTarget{t}
}

type Message struct{
	Target Target
	Content string
	From Target
}

func (m Message) To(t Target) Message {
	return m
}

func (m Message) Text(text string) Message {
	return m
}

func (m Message) Tag(t Target) Message {
	return m
}

func (m Message) SendFrom(t Target) {
}

type ListenFunc func(req Message)

func Global(id string, f ListenFunc) ListenFunc {
	return f
}

func Tagged(id string, f ListenFunc) ListenFunc {
	return f
}
