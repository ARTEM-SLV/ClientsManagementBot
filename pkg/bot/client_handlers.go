package bot

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tucnak/telebot"
)

// Структура для хранения заявок пользователей
var pendingRequests = make(map[int]*telebot.User)

// Обработчик для создания новой записи
func NewClientHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		log.Println("Получена команда /new")
		bot.Send(m.Sender, "Введите ваше имя для записи на услугу.")
		// Логика для создания записи клиента
	}
}

// Обработчик для просмотра прайс-листа
func ListServicesHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		log.Println("Получена команда /list_services")
		bot.Send(m.Sender, "Вот список доступных услуг:")
		// Логика для отображения списка услуг
	}
}

// Обработчик для просмотра расписания
func ScheduleHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		log.Println("Получена команда /schedule")
		bot.Send(m.Sender, "Вот расписание работы:")
		// Логика для отображения расписания
	}
}

// Обработчик команды /start для новых пользователей
func StartHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if isUserAuthorized(m.Sender.ID) {
			bot.Send(m.Sender, "Добро пожаловать обратно! Вы уже авторизованы.")
			return
		}

		// Сохраняем пользователя в список ожидания
		pendingRequests[m.Sender.ID] = m.Sender
		bot.Send(m.Sender, "Ваш запрос отправлен администратору для подтверждения. Ожидайте решения.")

		// Отправляем админу запрос на авторизацию пользователя
		adminUserID, _ := strconv.Atoi(adminID)
		bot.Send(&telebot.User{ID: adminUserID}, fmt.Sprintf("Новый пользователь: %s %s (ID: %d) запросил доступ.", m.Sender.FirstName, m.Sender.LastName, m.Sender.ID))

		// Создаем инлайн-кнопки
		acceptBtn := telebot.InlineButton{
			Unique: fmt.Sprintf("accept_user_%d", m.Sender.ID),
			Text:   "Принять",
		}
		rejectBtn := telebot.InlineButton{
			Unique: fmt.Sprintf("reject_user_%d", m.Sender.ID),
			Text:   "Отклонить",
		}

		// Создаем разметку для кнопок
		markup := &telebot.ReplyMarkup{}
		markup.Inline(
			markup.Row(acceptBtn, rejectBtn),
		)

		// Отправляем админу сообщение с кнопками
		bot.Send(&telebot.User{ID: adminUserID}, "Выберите действие:", markup)
	}
}

// Проверка, авторизован ли пользователь
func isUserAuthorized(userID int) bool {
	// Здесь должна быть логика проверки авторизации (например, проверка в базе данных)
	// Пока фиктивная проверка, всегда false для новых пользователей
	return false
}
