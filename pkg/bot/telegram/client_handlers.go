package telegram

import (
	"github.com/tucnak/telebot"
	"log"
)

func btnClientScheduleFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Вы выбрали: Расписание")
		// Здесь можно добавить логику для отображения списка услуг
	}
}

func btnClientServicesFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Вы выбрали: Выбор услуги")
		// Здесь можно добавить логику для отображения списка услуг
	}
}

func btnClientHelpFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Справка: Вы можете выбрать услугу или получить помощь. "+
			"Для подробностей обращайтесь к администратору.")
	}
}

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
