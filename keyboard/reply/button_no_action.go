package reply

import (
	"github.com/go-telegram/bot/models"
)

func (kb *ReplyKeyboard) ButtonNoAction(text string) *ReplyKeyboard {
	kb.markup[len(kb.markup)-1] = append(kb.markup[len(kb.markup)-1], models.KeyboardButton{
		Text: text,
	})

	return kb
}

func (kb *ReplyKeyboard) ButtonNoActionWithRequestContacts(text string) *ReplyKeyboard {
	kb.markup[len(kb.markup)-1] = append(kb.markup[len(kb.markup)-1], models.KeyboardButton{
		Text:           text,
		RequestContact: true,
	})

	return kb
}
