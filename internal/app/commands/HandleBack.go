package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ComRout) HandleBack(inputMessage *tgbotapi.Message) {

	c.PopState(inputMessage.Chat.ID)
	previousState := c.GetState(inputMessage.Chat.ID)
	if previousState == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Нет предыдущего раздела.")
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
		}
		return
	}
	if states, exists := c.userStates[inputMessage.Chat.ID]; exists && len(states) > 0 {
		lastState := states[len(states)-1]
		log.Printf("Help: ID чата: %d, текущее состояние: %s, полный стек: %v", inputMessage.Chat.ID, lastState, states)
	} else {
		log.Printf("Help: ID чата: %d, стек состояний пуст.", inputMessage.Chat.ID)
	}
	switch previousState {
	case "list":
		c.List(inputMessage)
	case "second":
		c.Vtoroi_kurs(inputMessage)
	case "3_sem":
		c.Tri_sem(inputMessage)
	case "start":
		c.Start(inputMessage)
	case "3_sem_2022":
		c.Tri_sem_2022(inputMessage)
	case "Kiir_telyak":
		c.Kiir_telyak(inputMessage)
	case "dopusk":
		c.Dopusk(inputMessage)
	case "kr1":
		c.Kr1(inputMessage)
	case "kr2":
		c.Kr2(inputMessage)
	case "kr3":
		c.Kr3(inputMessage)
	case "kr4":
		c.Kr4(inputMessage)
	default:
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Неизвестное состояние.")
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
		}
	}
}
