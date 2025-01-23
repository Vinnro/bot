package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ComRout) Vtoroi_kurs(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "second")
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Выберите нужные вам позиции: ")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("3_сем", "3_sem"),
				tgbotapi.NewInlineKeyboardButtonData("4_сем", "4_sem"),
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

func (c *ComRout) Tri_sem(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "3_sem")
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Выберите нужные вам позиции: ")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("3_сем_2022", "3_sem_2022"),
				tgbotapi.NewInlineKeyboardButtonData("3_сем_2023", "3_sem_2023"),
				tgbotapi.NewInlineKeyboardButtonData("3_сем_2021", "3_sem_2021"),
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

func (c *ComRout) Tri_sem_2022(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "3_sem_2022")
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Выберите нужные вам позиции: ")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("Кратные интегралы и ряды _ Преп_ Теляковский", "Kiir_Tel"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("Специальные главы линейной алгебры _ Преп_ Тюфлин", "SpecLinal_T"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("Физ.Практикум _ Препы_ Тюлюсов _) - Кумпан _", "FizPract"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("Физика (электричество и магнетизм) _ Преп_ Мичулис", "Phizika_Mich"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("Электротехника _ Преп_ Новожилов", "Altex_Nov"),
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

func (c *ComRout) Kiir_telyak(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "Kiir_telyak")
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Выберите нужные вам позиции: ")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("Допуск", "Dopusk"),
				tgbotapi.NewInlineKeyboardButtonData("Kr1", "Kr1"),
			},
			{
				tgbotapi.NewInlineKeyboardButtonData("Kr2", "Kr2"),
				tgbotapi.NewInlineKeyboardButtonData("Kr3", "Kr3"),
				tgbotapi.NewInlineKeyboardButtonData("Kr4", "Kr4"),
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

func (c *ComRout) Dopusk(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "dopusk")
	photo := tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FilePath("/home/kailo/Kom.webp"))
	photo.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
			},
		},
	}
	if _, err := c.bot.Send(photo); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func (c *ComRout) Kr1(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "kr1")
	photo := tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FilePath("/path/to/Kr1_image.webp"))
	if _, err := c.bot.Send(photo); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func (c *ComRout) Kr2(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "kr2")
	photo := tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FilePath("/path/to/Kr2_image.webp"))
	if _, err := c.bot.Send(photo); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func (c *ComRout) Kr3(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "kr3")
	photo := tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FilePath("/path/to/Kr3_image.webp"))
	if _, err := c.bot.Send(photo); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func (c *ComRout) Kr4(inputMessage *tgbotapi.Message) {
	c.pushState(inputMessage.Chat.ID, "kr4")
	photo := tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FilePath("/path/to/Kr4_image.webp"))
	if _, err := c.bot.Send(photo); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}
