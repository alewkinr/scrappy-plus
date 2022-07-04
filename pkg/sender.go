package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pkgz/notify"
)

// Sender — отправитель сообщений
type Sender struct {
	// ChannelID — идентификатор канала для отправки сообщений
	ChannelID string
	// token – токен для доступа к телеграмму
	token string
	// tg — клиент для телеграма
	tg *notify.Telegram
}

// MustInitSender — конструктор отправщика сообщений
func MustInitSender(channelID, token string) *Sender {
	tg, err := notify.NewTelegram(notify.TelegramParams{
		Token:   token,
		Timeout: time.Second * 10,
	})
	if err != nil {
		panic(fmt.Sprintf("init sender: %s", err))
	}

	return &Sender{
		ChannelID: "telegram:" + channelID + "?parseMode=HTML",
		token:     token,
		tg:        tg,
	}
}

// Send — отправляем сообщение
func (s *Sender) Send(ctx context.Context, msg string) error {
	return s.tg.Send(ctx, s.ChannelID, msg)
}
