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

		// Отправляем приветственное сообщение и отображаем меню
		var msg string
		if bot.IsAdmin(m.Sender.ID) {
			msg = bot.MessagesList.WelcomeAdmin
			menu.ReplyKeyboard = createBtnAdmin(b)
		} else {
			msg = fmt.Sprintf(bot.MessagesList.WelcomeClient, m.Sender.LastName, m.Sender.FirstName)
			menu.ReplyKeyboard = createBtnClient(b)
		}
		b.Send(m.Sender, msg, menu)
	}
}

func createBtnClient(b *telebot.Bot) [][]telebot.ReplyButton {
	// Создаем кнопки
	btnServices := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnServices}
	btnHelp := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnHelp}

	// Обработчики для кнопок
	b.Handle(&btnServices, btnServiceFunc)
	b.Handle(&btnHelp, btnHelpFunc)

	return [][]telebot.ReplyButton{
		{btnServices}, // "Выбор услуги"
		{btnHelp},     // "Справка"
	}
}

func createBtnAdmin(b *telebot.Bot) [][]telebot.ReplyButton {
	// Создаем кнопки
	btnSchedule := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnSchedule}
	btnServices := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnServices}
	btnReports := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnReports}
	btnHelp := telebot.ReplyButton{Text: bot.BtnTitlesList.BtnHelp}

	// Обработчики для кнопок
	b.Handle(&btnSchedule, btnScheduleFunc)
	b.Handle(&btnServices, btnServiceFunc)
	b.Handle(&btnReports, btnReportsFunc)
	b.Handle(&btnHelp, btnHelpFunc)

	return [][]telebot.ReplyButton{
		{btnSchedule}, // "Расписание"
		{btnServices}, // "Выбор услуги"
		{btnReports},  // "Отчеты"
		{btnHelp},     // "Справка"
	}
}

func btnScheduleFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Вы выбрали: Расписание")
		// Здесь можно добавить логику для отображения списка услуг
	}
}

func btnServiceFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Вы выбрали: Выбор услуги")
		// Здесь можно добавить логику для отображения списка услуг
	}
}
func btnReportsFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Вы выбрали: Отчеты")
		// Здесь можно добавить логику для отображения списка услуг
	}
}

func btnHelpFunc(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		b.Send(m.Sender, "Справка: Вы можете выбрать услугу или получить помощь. "+
			"Для подробностей обращайтесь к администратору.")
	}
}
