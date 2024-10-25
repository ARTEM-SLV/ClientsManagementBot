package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/tucnak/telebot"

	"ClientsManagementBot/pkg/bot"
	"ClientsManagementBot/pkg/bot/telegram"
)

func main() {
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	// Получаем токен для бота
	botToken := os.Getenv("CLIENTS_MANAGEMENT_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Токен бота не найден")
	}

	// Инициализация бота
	err = bot.InitBot()
	if err != nil {
		log.Fatal("Ошибка при инициализации бота:", err)
	}

	// Инициализация бота Телеграм
	botInstance, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("Ошибка при инициализации бота Телеграм:", err)
	}

	// Подключаем обработчик для команды /start
	botInstance.Handle("/start", telegram.StartHandler(botInstance))

	// Подключение других административных команд
	botInstance.Handle("/add_service", telegram.AddServiceHandler(botInstance))
	botInstance.Handle("/list_clients", telegram.ListClientsHandler(botInstance))
	botInstance.Handle("/schedule_update", telegram.ScheduleUpdateHandler(botInstance))

	log.Println("Бот запущен и готов к работе...")

	// Запуск бота
	botInstance.Start()
}
