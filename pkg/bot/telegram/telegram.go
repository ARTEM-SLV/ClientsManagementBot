package telegram

import (
	"fmt"
	"github.com/tucnak/telebot"

	"ClientsManagementBot/pkg/bot"
)

// Обработчик команды /start
func StartHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		// Создаем клавиатуру с кнопками
		menu := &telebot.ReplyMarkup{ResizeReplyKeyboard: true}
		//btnMenu := menu.InlineKeyboard

		// Отправляем приветственное сообщение и отображаем меню
		var msg string
		if bot.IsAdmin(m.Sender.ID) {
			msg = bot.MessagesList.WelcomeAdmin
			menu.InlineKeyboard = createBtnAdmin(b)
		} else {
			msg = fmt.Sprintf(bot.MessagesList.WelcomeClient, m.Sender.LastName, m.Sender.FirstName)
			menu.InlineKeyboard = createBtnClient(b)
		}
		b.Send(m.Sender, msg, menu)
	}
}

func createBtnClient(b *telebot.Bot) [][]telebot.InlineButton {
	// Создаем кнопки
	btnServices := telebot.InlineButton{
		Unique: "btn_services",
		Text:   bot.BtnTitlesList.BtnServices,
	}
	btnHelp := telebot.InlineButton{
		Unique: "btn_help",
		Text:   bot.BtnTitlesList.BtnHelp,
	}

	// Обработчики для кнопок
	b.Handle(&btnServices, btnServiceFunc)
	b.Handle(&btnHelp, btnHelpFunc)

	return [][]telebot.InlineButton{
		{btnServices}, // "Выбор услуги"
		{btnHelp},     // "Справка"
	}
}

func createBtnAdmin(b *telebot.Bot) [][]telebot.InlineButton {
	// Создаем кнопки
	btnSchedule := telebot.InlineButton{
		Unique: "btn_schedule",
		Text:   bot.BtnTitlesList.BtnSchedule,
	}
	btnServices := telebot.InlineButton{
		Unique: "btn_services",
		Text:   bot.BtnTitlesList.BtnServices,
	}
	btnReports := telebot.InlineButton{
		Unique: "btn_reports",
		Text:   bot.BtnTitlesList.BtnReports,
	}
	btnHelp := telebot.InlineButton{
		Unique: "btn_help",
		Text:   bot.BtnTitlesList.BtnHelp,
	}

	// Обработчики для кнопок
	b.Handle(&btnSchedule, btnScheduleFunc)
	b.Handle(&btnServices, btnServiceFunc)
	b.Handle(&btnReports, btnReportsFunc)
	b.Handle(&btnHelp, btnHelpFunc)

	return [][]telebot.InlineButton{
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
