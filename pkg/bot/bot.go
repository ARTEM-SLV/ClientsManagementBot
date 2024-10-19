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
	BtnServices string `json:"btn_services"`
	BtnHelp     string `json:"btn_help"`
}

var adminID string
var MessagesList Messages
var BtnTitlesList BtnTitles

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

// Проверяем, является ли пользователь администратором
func IsAdmin(userID int) bool {
	return adminID == strconv.Itoa(userID)
}
