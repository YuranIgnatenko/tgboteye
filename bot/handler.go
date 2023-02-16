package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var IgnoreList = []string{
	"Добавить",
	"Отмена",
	"Сохранить",
	"История",
	"Активность",
	"Остановить",
	"Остановить все",
	"Забыть", 
	"Забыто", 
	"Ошибка добавления", 
	"/start",
}
var PrefixListUser = "Цель: "
var RestartFlag = true
var RecoverFlag = false

func (b *Bot) HandlerButton(userText string) (string, tgbotapi.ReplyKeyboardMarkup) {
	var msgtext = "+"
	var keyboard = b.KeyboardCreate(b.MenuMain())

	if RestartFlag==true{
		data,_:=LoadTasks() 
		if len(data) > 0{
			msgtext="Найдены задачи\n" 
			for _, line:=range data{
				msgtext+=line+"\n" 
			} 
			RecoverFlag = true
			RestartFlag=false
			keyboard = b.KeyboardCreate(b.MenuMain())
			RecoverFlag=false
			return msgtext, keyboard
		} 
	} 

	switch userText {

	case "Восстановить": 
		msgtext = "восстановлено" 
		b.RecoveryTasks()
		keyboard = b.KeyboardCreate(b.MenuMain())
		//RecoverFlag =false

	case "Добавить":
		msgtext = "username/phone: @name / +79998887766"
		keyboard = b.KeyboardCreate(b.MenuAddUser())

	case "Отмена":
		keyboard = b.KeyboardCreate(b.MenuMain())

	case "Сохранить":
		msgtext = "Добавлено: " + b.LastMessage
		err := b.AddTask(b.LastMessage)
		if err != nil {
			switch err{
				case ErrorAddTask: 
					msgtext="Ошибка добавления" 
				default: 
					msgtext = err.Error()
			} 
		}
		b.SaveUserRequest(b.LastMessage)
		keyboard = b.KeyboardCreate(b.MenuMain())

	case "История":
		if b.ReadUserRequest() == "" {
			msgtext = "История запросов пуста"
			keyboard = b.KeyboardCreate(b.MenuMain())
		} else {
			msgtext = "Прошлые цели: \n" + b.ReadUserRequest()
			keyboard = b.KeyboardCreate(b.MenuHistory())
		}

	case "Удалить":
		msgtext = "История запросов удалена"
		b.DeleteUserRequest()
		keyboard = b.KeyboardCreate(b.MenuMain())

	case "Активность":
		data, err := ioutil.ReadFile(fmt.Sprintf("file_eye_%s.log", b.LastMessage))
		if err != nil {
			log.Println(err)
		}
		msgtext = "Запись активности:\n" + string(data)
		keyboard = b.KeyboardCreate(b.MenuMain())

	case "Забыть": 
		LineGet(fmt.Sprintf("rm file_eye_%s.log", b. LastMessage)) 
		msgtext = "Забыто" 

	case "Остановить":
		msgtext = "Остановлен: " + b.LastMessage
		err := b.AbortTask(b.LastMessage)
		if err != nil {
			log.Println(err)
		}
		keyboard = b.KeyboardCreate(b.MenuMain())

	case "Остановить все":
		users_list, err := b.UsersList()
		if err != nil {
			msgtext = "Добавь цели"
		} else {
			msgtext = fmt.Sprintf("Отменены:\n%v", users_list)
			err = b.AbortTaskAll()
			if err != nil {
				log.Println(err)
			}
		}
		keyboard = b.KeyboardCreate(b.MenuMain())

	default:
		for _, com := range IgnoreList {
			if userText == com {
				return msgtext, keyboard
			}
		}
		if strings.HasPrefix(userText, PrefixListUser) {
			b.LastMessage = strings.ReplaceAll(userText, PrefixListUser, "")
			keyboard = b.KeyboardCreate(b.MenuSelectUser())
		} else {
			msgtext = "Сохранить/Забыть" + userText + " ?"
			keyboard = b.KeyboardCreate(b.MenuAddUser())
			b.LastMessage = userText
		}
	}
	return msgtext, keyboard

}
