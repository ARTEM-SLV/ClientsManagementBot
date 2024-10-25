package main

import (
	"ClientsManagementBot/pkg/database"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/tucnak/telebot"

	"ClientsManagementBot/pkg/bot"
	"ClientsManagementBot/pkg/bot/telegram"
)

func main() {
	InitLogger()

	dbType := "sqlite" // "sqlite" или "postgres"
	db, err := database.NewDatabase(dbType)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	if err := db.Init(); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	// Загружаем переменные окружения
	err = godotenv.Load()
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

func InitLogger() {
	// Открываем файл для логов (или создаем, если его нет)
	file, err := os.OpenFile("bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла для логов: %v", err)
	}

	// Настройка логгера для записи в файл
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
