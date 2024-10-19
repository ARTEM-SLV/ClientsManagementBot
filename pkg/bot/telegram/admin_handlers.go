package telegram

import (
	"github.com/tucnak/telebot"
	"log"
	"os"
	"strconv"
)

// Администраторский Telegram ID (можно хранить в .env)
var adminID = os.Getenv("ADMIN_ID")

// Проверяем, является ли пользователь администратором
func isAdmin(userID int) bool {
	return adminID == strconv.Itoa(userID)
}

// Обработчик для добавления новой услуги (только для администратора)
func AddServiceHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if !isAdmin(m.Sender.ID) {
			b.Send(m.Sender, "Эта команда доступна только администратору.")
			return
		}
		log.Println("Получена команда /add_service")
		b.Send(m.Sender, "Введите описание новой услуги.")
		// Логика для добавления услуги
	}
}

// Обработчик для просмотра всех клиентов (только для администратора)
func ListClientsHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if !isAdmin(m.Sender.ID) {
			b.Send(m.Sender, "Эта команда доступна только администратору.")
			return
		}
		log.Println("Получена команда /list_clients")
		b.Send(m.Sender, "Вот список всех записей клиентов:")
		// Логика для просмотра записей всех клиентов
	}
}

// Обработчик для обновления расписания (только для администратора)
func ScheduleUpdateHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if !isAdmin(m.Sender.ID) {
			b.Send(m.Sender, "Эта команда доступна только администратору.")
			return
		}
		log.Println("Получена команда /schedule_update")
		b.Send(m.Sender, "Введите новое расписание.")
		// Логика для обновления расписания
	}
}
