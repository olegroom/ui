package reply

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (kb *ReplyKeyboard) ButtonNoAction(text string, b *bot.Bot, matchType bot.MatchType) *ReplyKeyboard {
	kb.markup[len(kb.markup)-1] = append(kb.markup[len(kb.markup)-1], models.KeyboardButton{
		Text: text,
	})

	return kb
}
