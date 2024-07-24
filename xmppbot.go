package xmppbot

import (
	"context"
	"fmt"
	"log" // @todo: slog?
	"runtime"

	"mellium.im/xmpp"
	"mellium.im/xmpp/jid"
	"mellium.im/sasl"
	"mellium.im/xmpp/muc"
	"mellium.im/xmpp/mux"
)

func New(ctx context.Context) *Bot {
	c := &muc.Client{}
	m := mux.New(
		"",
		muc.HandleClient(c),
	)
	b := &Bot{
		Ctx: ctx,
		Mux: m,
		MucClient: c,
	}
	return b
}

type Bot struct{
	Target
	Ctx context.Context
	Session *xmpp.Session
	Mux *mux.ServeMux
	Channel *muc.Channel
	MucClient *muc.Client
	Error error
}

func (b *Bot) reportError(err error) bool {
	if err != nil {
		pc, f, l, _ := runtime.Caller(1)
		b.Error = fmt.Errorf("%s[%s:%d]: %v", runtime.FuncForPC(pc).Name(), f, l, err)
		log.Printf("xmppbot: %s", b.Error)
		return false
	}
	return true
}

func (b *Bot) Login(user, pass string) *Bot {
	id, err := jid.Parse(user)
	if !b.reportError(err) {
		return b
	}
	session, err := xmpp.DialClientSession(
		b.Ctx,
		id,
		xmpp.StartTLS(nil),
		xmpp.SASL("", pass, sasl.ScramSha256Plus, sasl.ScramSha256, sasl.ScramSha1Plus, sasl.ScramSha1, sasl.Plain),
		xmpp.BindResource(),
	)
	if !b.reportError(err) {
		return b
	}
	b.Session = session
	go b.Session.Serve(b.Mux)
	return b
}

func (b *Bot) Join(t RoomTarget) *Bot {
	channel, err := b.MucClient.Join(b.Ctx, t.Jid, b.Session, t.Opts...)
	b.reportError(err)
	b.Channel = channel
	return b
}

func (b *Bot) SendMessage(msg Message) {
}

func (b *Bot) ListenToCommand(f ListenFunc) {
}

func (b *Bot) SendTextMessage(text string) {
}

type (
	Target struct{
		Jid jid.JID
	}
	RoomTarget struct{
		Target
		Opts []muc.Option
	}
	ContactTarget struct{
		Target
	}
)

func (t Target) ToTarget() Target {
	return t
}

func (t Target) Room(room string, opts ...muc.Option) RoomTarget {
	id := jid.MustParse(room)
	t.Jid = id
	return RoomTarget{
		Target: t,
		Opts: opts,
	}
}

func (t Target) Contact(contact string) ContactTarget {
	id := jid.MustParse(contact)
	t.Jid = id
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
