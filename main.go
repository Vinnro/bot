package main

import (
	"MephiBot/internal/app/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	ChatID := int64(5584877461)
	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	commander := commands.NewComRout(bot)

	for update := range updates {
		if update.Message != nil {
			prev := commander.GetState(update.Message.Chat.ID)
			if prev == "contact" {
				commander.PopState(update.Message.Chat.ID)
				msg := tgbotapi.NewMessage(ChatID, update.Message.Text+"\n"+"От пользователя: "+"@"+update.Message.From.UserName)
				bot.Send(msg)
				msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше сообщение отправлено. Вернитесь на главную страницу /start")
				bot.Send(msg2)
				continue
			}
			switch update.Message.Text {
			case "/start":
				commander.Start(update.Message)
			case "/help":
				commander.Help(update.Message)
			case "/list":
				commander.List(update.Message)
			case "/contact":
				commander.ContactHandler(update.Message)
			default:
				commander.Default(update.Message)
			}
		}

		if update.CallbackQuery != nil && update.CallbackQuery.Message != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			if _, err := bot.Request(callback); err != nil {
				log.Printf("Ошибка при отправке CallbackQuery: %v", err)
			}
			switch update.CallbackQuery.Data {
			case "contact":
				commander.ContactHandler(update.CallbackQuery.Message)
			case "started":
				commander.Start(update.CallbackQuery.Message)
			case "list":
				commander.List(update.CallbackQuery.Message)
			case "help":
				commander.Help(update.CallbackQuery.Message)
			case "back":
				commander.HandleBack(update.CallbackQuery.Message)
			case "second":
				commander.Vtoroi_kurs(update.CallbackQuery.Message)
			case "3_sem":
				commander.Tri_sem(update.CallbackQuery.Message)
			case "3_sem_2022":
				commander.Tri_sem_2022(update.CallbackQuery.Message)
			case "Kiir_Tel":
				commander.Kiir_telyak(update.CallbackQuery.Message)
				//case "Dopusk":
				//	commander.Dopusk(update.CallbackQuery.Message)
			}
		}
	}
}
