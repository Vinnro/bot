package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *ComRout) Start(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "start")
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Добро пожаловать в архив ИИКС мяу мяу МИФИ\n"+
		"Команды, которые тебе доступны:\n"+
		"👉 помощь - список всех доступных команд.\n"+
		"👉 список - показать категории материалов.\n"+
		"👉 поиск - поиск по материалам.\n"+
		"👉 связь - связаться с разработчиком.\n"+
		"И да спасёт нас бог\n")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("Помощь", "help"),
				tgbotapi.NewInlineKeyboardButtonData("Список", "list"),
				tgbotapi.NewInlineKeyboardButtonData("Поиск", "search"),
				tgbotapi.NewInlineKeyboardButtonData("Связь", "contact"),
			},
		},
	}
	c.bot.Send(msg)
}
