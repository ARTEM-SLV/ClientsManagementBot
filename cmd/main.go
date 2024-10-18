package main

import (
	"ClientsManagementBot/pkg/bot"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/tucnak/telebot"
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
	botInstance, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("Ошибка при инициализации бота:", err)
	}

	// Подключение клиентских команд
	botInstance.Handle("/start", bot.StartHandler(botInstance))

	// Подключение административных команд
	botInstance.Handle("/add_service", bot.AddServiceHandler(botInstance))
	botInstance.Handle("/list_clients", bot.ListClientsHandler(botInstance))
	botInstance.Handle("/schedule_update", bot.ScheduleUpdateHandler(botInstance))
	botInstance.Handle("/list_period", bot.ListClientsByPeriodHandler(botInstance))

	// Обработчик для колбеков
	botInstance.Handle(telebot.OnCallback, bot.UserRequestHandler(botInstance))

	log.Println("Бот запущен и готов к работе...")

	// Запуск бота
	botInstance.Start()
}
