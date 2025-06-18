package reply

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (kb *ReplyKeyboard) ButtonWithRequestContacts(text string, b *bot.Bot, matchType bot.MatchType, handler bot.HandlerFunc) *ReplyKeyboard {
	kb.markup[len(kb.markup)-1] = append(kb.markup[len(kb.markup)-1], models.KeyboardButton{
		Text:           text,
		RequestContact: true,
	})

	if handler != nil {
		b.RegisterHandler(bot.HandlerTypeMessageText, text, matchType, handler)
	}

	return kb
}
