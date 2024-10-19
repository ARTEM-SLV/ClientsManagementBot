package telegram

import (
	"ClientsManagementBot/pkg/bot"
	"fmt"
	"github.com/tucnak/telebot"
)

// Обработчик команды /start
func StartHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		// Создаем клавиатуру с кнопками
		menu := &telebot.ReplyMarkup{ResizeReplyKeyboard: true}

		// Создаем кнопки
		btnServices := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnServices}
		btnHelp := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnHelp}

		// Устанавливаем кнопки в меню (одна кнопка на строке)
		menu.ReplyKeyboard = [][]telebot.ReplyButton{
			{btnServices}, // Первая строка: одна кнопка "Выбор услуги"
			{btnHelp},     // Вторая строка: одна кнопка "Help"
		}

		// Отправляем приветственное сообщение и отображаем меню
		var msg string
		if bot.IsAdmin(m.Sender.ID) {
			msg = bot.MessagesList.WelcomeAdmin
		} else {
			msg = fmt.Sprintf(bot.MessagesList.WelcomeClient, m.Sender.LastName, m.Sender.FirstName)
		}
		b.Send(m.Sender, msg, menu)

		// Обработчики для кнопок
		b.Handle(&btnServices, btnServiceFunc)

		b.Handle(&btnHelp, btnHelpFunc)
	}
}

func btnServiceFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Вы выбрали: Выбор услуги")
		// Здесь можно добавить логику для отображения списка услуг
	}
}

func btnHelpFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Справка: Вы можете выбрать услугу или получить помощь. "+
			"Для подробностей обращайтесь к администратору.")
	}
}
