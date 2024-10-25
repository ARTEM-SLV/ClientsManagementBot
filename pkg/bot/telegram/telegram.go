package telegram

import (
	"fmt"
	"log"

	"github.com/tucnak/telebot"

	"ClientsManagementBot/pkg/bot"
)

// Обработчик команды /start
func StartHandler(b *telebot.Bot) func(*telebot.Message) {
	return func(m *telebot.Message) {
		// Создаем клавиатуру с кнопками
		menu := &telebot.ReplyMarkup{ResizeReplyKeyboard: true}

		// Отправляем приветственное сообщение и отображаем меню
		var msg string
		if bot.IsMaster(m.Sender.ID) {
			msg = fmt.Sprintf(bot.MessagesList.WelcomeAdmin, m.Sender.LastName, m.Sender.FirstName)
			menu.InlineKeyboard = createBtnMaster(b)
		} else {
			msg = fmt.Sprintf(bot.MessagesList.WelcomeClient, m.Sender.LastName, m.Sender.FirstName)
			menu.InlineKeyboard = createBtnClient(b)
		}
		_, err := b.Send(m.Sender, msg, menu)
		if err != nil {
			log.Println("Не удалось отправить приветственное сообщение:", err)
		}
	}
}

func createBtnClient(b *telebot.Bot) [][]telebot.InlineButton {
	// Создаем кнопки
	btnSchedule := telebot.InlineButton{
		Unique: "btn_schedule",
		Text:   bot.BtnTitlesList.BtnSchedule,
	}
	btnServices := telebot.InlineButton{
		Unique: "btn_services",
		Text:   bot.BtnTitlesList.BtnServices,
	}
	btnHelp := telebot.InlineButton{
		Unique: "btn_help",
		Text:   bot.BtnTitlesList.BtnHelp,
	}

	// Обработчики для кнопок
	b.Handle(&btnSchedule, btnClientScheduleFunc)
	b.Handle(&btnServices, btnClientServicesFunc)
	b.Handle(&btnHelp, btnClientHelpFunc)

	return [][]telebot.InlineButton{
		{btnSchedule}, // "Расписание"
		{btnServices}, // "Выбор услуги"
		{btnHelp},     // "Справка"
	}
}

func createBtnMaster(b *telebot.Bot) [][]telebot.InlineButton {
	// Создаем кнопки
	btnSchedule := telebot.InlineButton{
		Unique: "btn_schedule",
		Text:   bot.BtnTitlesList.BtnSchedule,
	}
	btnReports := telebot.InlineButton{
		Unique: "btn_reports",
		Text:   bot.BtnTitlesList.BtnReports,
	}
	btnSettings := telebot.InlineButton{
		Unique: "btn_settings",
		Text:   bot.BtnTitlesList.BtnSettings,
	}
	btnHelp := telebot.InlineButton{
		Unique: "btn_help",
		Text:   bot.BtnTitlesList.BtnHelp,
	}

	// Обработчики для кнопок
	b.Handle(&btnSchedule, btnMasterScheduleFunc)
	b.Handle(&btnReports, btnReportsFunc)
	b.Handle(&btnSettings, btnSettingsFunc)
	b.Handle(&btnHelp, btnMasterHelpFunc)

	return [][]telebot.InlineButton{
		{btnSchedule}, // "Расписание"
		{btnReports},  // "Отчеты"
		{btnSettings}, // "Настройки"
		{btnHelp},     // "Справка"
	}
}
