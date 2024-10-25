package bot

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

type Messages struct {
	WelcomeAdmin  string `json:"welcome_admin"`
	WelcomeClient string `json:"welcome_client"`
	ChooseAction  string `json:"choose_action"`
	HelpText      string `json:"help_text"`
}

type BtnTitles struct {
	BtnMenu     string `json:"btn_menu"`
	BtnBack     string `json:"btn_back"`
	BtnHelp     string `json:"btn_help"`
	BtnServices string `json:"btn_services"`
	BtnSchedule string `json:"btn_schedule"`
	BtnReports  string `json:"btn_reports"`
	BtnSettings string `json:"btn_settings"`
}

var adminID string
var MessagesList Messages
var BtnTitlesList BtnTitles

// Проверяем, является ли пользователь администратором
func IsMaster(userID int) bool {
	return adminID == strconv.Itoa(userID)
}

func InitBot() error {
	// Получаем ИД администратора
	adminID = os.Getenv("ADMIN_ID")

	// Загружаем сообщения бота
	err := loadMessages("./configs/messages.json")
	if err != nil {
		return err
	}

	// Загружаем заголовки кнопок
	err = loadBtnTitles("./configs/button_titles.json")
	if err != nil {
		return err
	}

	return nil
}

func loadMessages(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	json.Unmarshal(byteValue, &MessagesList)

	return nil
}

func loadBtnTitles(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	json.Unmarshal(byteValue, &BtnTitlesList)

	return nil
}
