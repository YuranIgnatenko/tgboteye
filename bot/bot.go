package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) KeyboardCreate(data []string) tgbotapi.ReplyKeyboardMarkup {
	var kk = make([][]tgbotapi.KeyboardButton, len(data))
	for i, value := range data {
		k := make([]tgbotapi.KeyboardButton, 1)
		k[0] = tgbotapi.NewKeyboardButton(value)
		kk[i] = k
	}
	return tgbotapi.NewReplyKeyboard(kk...)
}

func StyleStatus(user string) string {
	return PrefixListUser + user //"🟢" +
}

type Bot struct {
	Token       string
	MsgBot      string
	Timeout     int
	Tasks       []*Task
	Keyboard    tgbotapi.ReplyKeyboardMarkup
	LastMessage string
}

func New(token string) *Bot {
	return &Bot{
		Token:       token,
		MsgBot:      "complete",
		Timeout:     60,
		Tasks:       make([]*Task, 0),
		Keyboard:    tgbotapi.ReplyKeyboardMarkup{},
		LastMessage: "",
	}
}

func (b *Bot) Start() {
	bot, err := tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		panic(err)
	}

	b.Keyboard = b.KeyboardCreate(b.MenuMain())

	u := tgbotapi.NewUpdate(0)
	u.Timeout = b.Timeout
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		} else {

			newText, keyboard := b.HandlerButton(update.Message.Text)
			b.Keyboard = keyboard
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, newText)
			// fmt.Printf("%#v\n", b.Keyboard)
			msg.ReplyMarkup = b.Keyboard
			bot.Send(msg)
		}
	}
}

func (b *Bot) AddTask(targetUser string) error {
	for _, t := range b.Tasks {
		if t.TargetUser == targetUser {
			return ErrorAddTask
		}
	}

	b.Tasks = append(b.Tasks, NewTask(targetUser))
	T := b.Tasks[len(b.Tasks)-1]
	T.TargetUser = targetUser
	go T.Start()
	return nil
}

func (b *Bot) AbortTask(targetUser string) error {
	if len(b.Tasks) == 0 {
		return ErrorAbortTask
	}
	var part1 = make([]*Task, 0)
	var part2 = make([]*Task, 0)
	var ind_task int

	for ind, t := range b.Tasks {
		if t.TargetUser == targetUser {
			part1 = b.Tasks[:ind]
			part2 = b.Tasks[ind+1:]
			ind_task = ind
		}
	}

	go b.Tasks[ind_task].Abort()

	b.Tasks = make([]*Task, 0)
	b.Tasks = append(b.Tasks, part1[:]...)
	b.Tasks = append(b.Tasks, part2[:]...)
	return nil
}

func (b *Bot) AbortTaskAll() error {
	if len(b.Tasks) == 0 {
		return ErrorAbortTask
	}
	for _, t := range b.Tasks {
		b.AbortTask(t.TargetUser)
	}
	return nil
}

func (b *Bot) StatusTask(targetUser string) (bool, error) {
	for _, t := range b.Tasks {
		if t.TargetUser == targetUser {
			return t.Status, nil
		}
	}
	return false, ErrorStatusTask
}

func (b *Bot) StatusAllTask() (string, error) {
	statusall := ""
	for _, t := range b.Tasks {
		statusall += fmt.Sprintf("-> %v = %v\n", t.TargetUser, t.Status)
	}

	if len(statusall) == 0 {
		return "", ErrorStatusAllTask
	}

	return statusall, nil
}

func (b *Bot) UsersList() (string, error) {
	statusall := ""
	for _, t := range b.Tasks {
		statusall += fmt.Sprintf("%v\n", t.TargetUser)
	}

	if len(statusall) == 0 {
		return "", ErrorListUsers
	}

	return statusall, nil
}

func (b *Bot) ReadUserRequest() string {
	data, err := ioutil.ReadFile("history_user_request")
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(data)
}

func (b *Bot) SaveUserRequest(user string) error {
	old := b.ReadUserRequest()
	data := strings.Split(old, "\n")
	var lastuser string
	if len(data) > 1 {
		lastuser = data[len(data)-1]
	} else {
		lastuser = data[0]
	}
	if strings.TrimSpace(lastuser) != strings.TrimSpace(user) {
		return ioutil.WriteFile("history_user_request", []byte(user+"\n"+old), 0755)
	}
	return nil //correct todo

}

func (b *Bot) DeleteUserRequest() error {
	return ioutil.WriteFile("history_user_request", []byte(""), 0755)
}

func (b *Bot) RecoveryTasks() {
	tasks, _ := LoadTasks()
	for _, user := range tasks {
		b.AddTask(user)
	}
}
