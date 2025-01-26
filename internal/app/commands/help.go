package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ComRout) Help(inputMessage *tgbotapi.Message) {
	c.PushState(inputMessage.Chat.ID, "helper")
	if states, exists := c.userStates[inputMessage.Chat.ID]; exists {
		log.Printf("Help: ID чата: %d, текущий стек состояний: %v", inputMessage.Chat.ID, states)
	} else {
		log.Printf("Help: ID чата: %d, стек состояний пуст или не существует.", inputMessage.Chat.ID)
	}
	if c.bot == nil {
		log.Panic("Ошибка: bot не инициализирован")
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Список команд:\n"+
			"👉 Список команд - /help \n"+
			"👉 Список разделов - /list\n"+
			"👉 Главная страница - /start\n")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
			},
		},
	}
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}
