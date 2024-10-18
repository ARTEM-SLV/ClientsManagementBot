package bot

import (
	"fmt"
	"github.com/tucnak/telebot"
	"log"
	"os"
	"strconv"
	"time"
)

// Администраторский Telegram ID (можно хранить в .env)
var adminID = os.Getenv("ADMIN_ID")

// Проверяем, является ли пользователь администратором
func isAdmin(userID int) bool {
	return adminID == strconv.Itoa(userID)
}

// Обработчик для добавления новой услуги (только для администратора)
func AddServiceHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if !isAdmin(m.Sender.ID) {
			bot.Send(m.Sender, "Эта команда доступна только администратору.")
			return
		}
		log.Println("Получена команда /add_service")
		bot.Send(m.Sender, "Введите описание новой услуги.")
		// Логика для добавления услуги
	}
}

// Обработчик для просмотра всех клиентов (только для администратора)
func ListClientsHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if !isAdmin(m.Sender.ID) {
			bot.Send(m.Sender, "Эта команда доступна только администратору.")
			return
		}
		log.Println("Получена команда /list_clients")
		bot.Send(m.Sender, "Вот список всех записей клиентов:")
		// Логика для просмотра записей всех клиентов
	}
}

// Обработчик для обновления расписания (только для администратора)
func ScheduleUpdateHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if !isAdmin(m.Sender.ID) {
			bot.Send(m.Sender, "Эта команда доступна только администратору.")
			return
		}
		log.Println("Получена команда /schedule_update")
		bot.Send(m.Sender, "Введите новое расписание.")
		// Логика для обновления расписания
	}
}

// Обработчик для просмотра записей за любой период (только для администратора)
func ListClientsByPeriodHandler(bot *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		if !isAdmin(m.Sender.ID) {
			bot.Send(m.Sender, "Эта команда доступна только администратору.")
			return
		}

		// Просим ввести начальную дату
		bot.Send(m.Sender, "Введите начальную дату в формате YYYY-MM-DD:")
		bot.Handle(telebot.OnText, func(dateMsg *telebot.Message) {
			startDate, err := time.Parse("2006-01-02", dateMsg.Text)
			if err != nil {
				bot.Send(dateMsg.Sender, "Неправильный формат даты. Пожалуйста, введите в формате YYYY-MM-DD.")
				return
			}

			// Просим ввести конечную дату
			bot.Send(dateMsg.Sender, "Введите конечную дату в формате YYYY-MM-DD:")
			bot.Handle(telebot.OnText, func(endDateMsg *telebot.Message) {
				endDate, err := time.Parse("2006-01-02", endDateMsg.Text)
				if err != nil {
					bot.Send(endDateMsg.Sender, "Неправильный формат даты. Пожалуйста, введите в формате YYYY-MM-DD.")
					return
				}

				if endDate.Before(startDate) {
					bot.Send(endDateMsg.Sender, "Конечная дата не может быть раньше начальной.")
					return
				}

				// Логика поиска записей за указанный период
				log.Printf("Записи клиентов с %s по %s\n", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

				// Здесь нужно добавить логику поиска в базе данных
				records := getRecordsByPeriod(startDate, endDate) // Вызов функции, которая ищет записи в БД

				if len(records) == 0 {
					bot.Send(endDateMsg.Sender, fmt.Sprintf("Записи с %s по %s не найдены.", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")))
				} else {
					// Выводим записи клиентам
					bot.Send(endDateMsg.Sender, formatRecords(records))
				}
			})
		})
	}
}

// Вспомогательная функция для форматирования записей клиентов
func formatRecords(records []string) string {
	result := "Записи клиентов за указанный период:\n"
	for _, record := range records {
		result += record + "\n"
	}
	return result
}

// Временная функция для имитации получения записей из БД
func getRecordsByPeriod(startDate, endDate time.Time) []string {
	// Здесь нужно сделать запрос к базе данных, пока пример с фиктивными данными
	return []string{
		"Клиент: Иван Иванов, Услуга: Стрижка, Дата: 2024-10-10",
		"Клиент: Петр Петров, Услуга: Массаж, Дата: 2024-10-12",
	}
}

// Обработчик для колбеков "Принять" и "Отклонить"
func UserRequestHandler(bot *telebot.Bot) func(c *telebot.Callback) {
	return func(c *telebot.Callback) {
		userID := extractUserIDFromCallbackData(c.Data)
		pendingUser := pendingRequests[userID]
		if pendingUser == nil {
			bot.Respond(c, &telebot.CallbackResponse{Text: "Заявка не найдена или уже обработана."})
			return
		}

		switch {
		case c.Data[:7] == "accept_":
			log.Printf("Администратор принял пользователя с ID: %d", userID)
			authorizeUser(pendingUser)
			bot.Send(pendingUser, "Ваш запрос принят! Добро пожаловать в систему.")
			delete(pendingRequests, userID)
			bot.Respond(c, &telebot.CallbackResponse{Text: "Пользователь успешно принят."})

		case c.Data[:7] == "reject_":
			log.Printf("Администратор отклонил пользователя с ID: %d", userID)
			bot.Send(pendingUser, "К сожалению, ваш запрос на доступ был отклонен.")
			delete(pendingRequests, userID)
			bot.Respond(c, &telebot.CallbackResponse{Text: "Пользователь отклонен."})
		}
	}
}

// Обработчик кнопки "Принять"
func AcceptUserHandler(bot *telebot.Bot) func(c *telebot.Callback) {
	return func(c *telebot.Callback) {
		userID := c.Sender.ID
		log.Printf("Администратор принял пользователя с ID: %d", userID)

		// Получаем информацию о пользователе из заявки
		pendingUser := pendingRequests[userID]
		if pendingUser == nil {
			bot.Respond(c, &telebot.CallbackResponse{Text: "Заявка не найдена или уже обработана."})
			return
		}

		// Добавляем пользователя в список авторизованных (в БД)
		authorizeUser(pendingUser)

		// Сообщаем пользователю о принятии
		bot.Send(pendingUser, "Ваш запрос принят! Добро пожаловать в систему.")

		// Удаляем заявку из ожидания
		delete(pendingRequests, userID)

		// Уведомляем администратора
		bot.Respond(c, &telebot.CallbackResponse{Text: "Пользователь успешно принят."})
	}
}

// Обработчик кнопки "Отклонить"
func RejectUserHandler(bot *telebot.Bot) func(c *telebot.Callback) {
	return func(c *telebot.Callback) {
		userID := c.Sender.ID
		log.Printf("Администратор отклонил пользователя с ID: %d", userID)

		// Получаем информацию о пользователе из заявки
		pendingUser := pendingRequests[userID]
		if pendingUser == nil {
			bot.Respond(c, &telebot.CallbackResponse{Text: "Заявка не найдена или уже обработана."})
			return
		}

		// Сообщаем пользователю об отказе
		bot.Send(pendingUser, "К сожалению, ваш запрос на доступ был отклонен.")

		// Удаляем заявку из ожидания
		delete(pendingRequests, userID)

		// Уведомляем администратора
		bot.Respond(c, &telebot.CallbackResponse{Text: "Пользователь отклонен."})
	}
}

// Вспомогательная функция для извлечения userID из данных колбека
func extractUserIDFromCallbackData(callbackData string) int {
	var userID int
	fmt.Sscanf(callbackData, "user_%d", &userID)
	return userID
}

// Временная функция для имитации добавления пользователя в БД
func authorizeUser(user *telebot.User) {
	// Здесь необходимо реализовать логику добавления пользователя в базу данных
	log.Printf("Пользователь %s (ID: %d) добавлен в авторизованные.", user.FirstName, user.ID)
}
