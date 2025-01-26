package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ComRout) List(inputMessage *tgbotapi.Message) {
	c.PushState(inputMessage.Chat.ID, "list")
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Выберите нужные вам позиции: ")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("1_КУРС_НАБОР_2023", "first2023"),
				tgbotapi.NewInlineKeyboardButtonData("1_КУРС_НАБОР_2022", "first2024"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("2_КУРС_НАБОР_2023", "second"),
				tgbotapi.NewInlineKeyboardButtonData("АРХИВЫ_МАТАНА", "аrchiv"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("КРУТОЙ_АРХИВ", "coolarchiv"),
				tgbotapi.NewInlineKeyboardButtonData("ОТЗЫВЫ_О_ПРЕПОДАХ", "rec"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
			},
		},
	}
	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}
