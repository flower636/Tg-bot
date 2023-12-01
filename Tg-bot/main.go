package main

import (
	"flag"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var (
	// глобальная переменная, в которой храним токен
	telegramBotToken string
)

func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {

	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настройка обработчик сообщений
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil { // игнорировать обновления, не являющиеся сообщениями
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Отвечаем на приветствие
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Добро пожаловать! Я бот Расписания"
			case "help":
				msg.Text = "Здесь будут прописаны команды навигации по боту"
			default:
				msg.Text = "Неизвестная команда"
			}
			bot.Send(msg)
		}
	}
}
