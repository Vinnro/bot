package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *ComRout) ContactHandler(inputMessage *tgbotapi.Message) {
	c.PushState(inputMessage.Chat.ID, "contact")
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Разработчик: @AtroxR \n Введите сообщение в этом чате: ")
	c.bot.Send(msg)
}
